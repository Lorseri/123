#include "MemTable.h"
#include "utils/Log.h"


namespace zilliz {
namespace milvus {
namespace engine {

MemTable::MemTable(const std::string &table_id,
                   const meta::MetaPtr &meta,
                   const Options &options) :
    table_id_(table_id),
    meta_(meta),
    options_(options) {

}

Status MemTable::Add(VectorSourcePtr &source, IDNumbers &vector_ids) {

    while (!source->AllAdded()) {

        MemTableFilePtr current_mem_table_file;
        if (!mem_table_file_list_.empty()) {
            current_mem_table_file = mem_table_file_list_.back();
        }

        Status status;
        if (mem_table_file_list_.empty() || current_mem_table_file->IsFull()) {
            MemTableFilePtr new_mem_table_file = std::make_shared<MemTableFile>(table_id_, meta_, options_);
            status = new_mem_table_file->Add(source, vector_ids);
            if (status.ok()) {
                mem_table_file_list_.emplace_back(new_mem_table_file);
            }
        } else {
            status = current_mem_table_file->Add(source, vector_ids);
        }

        if (!status.ok()) {
            std::string err_msg = "MemTable::Add failed: " + status.ToString();
            ENGINE_LOG_ERROR << err_msg;
            return Status(DB_ERROR, err_msg);
        }
    }
    return Status::OK();
}

void MemTable::GetCurrentMemTableFile(MemTableFilePtr &mem_table_file) {
    mem_table_file = mem_table_file_list_.back();
}

size_t MemTable::GetTableFileCount() {
    return mem_table_file_list_.size();
}

Status MemTable::Serialize() {
    for (auto mem_table_file = mem_table_file_list_.begin(); mem_table_file != mem_table_file_list_.end();) {
        auto status = (*mem_table_file)->Serialize();
        if (!status.ok()) {
            std::string err_msg = "MemTable::Serialize failed: " + status.ToString();
            ENGINE_LOG_ERROR << err_msg;
            return Status(DB_ERROR, err_msg);
        }
        std::lock_guard<std::mutex> lock(mutex_);
        mem_table_file = mem_table_file_list_.erase(mem_table_file);
    }
    return Status::OK();
}

bool MemTable::Empty() {
    return mem_table_file_list_.empty();
}

const std::string &MemTable::GetTableId() const {
    return table_id_;
}

size_t MemTable::GetCurrentMem() {
    std::lock_guard<std::mutex> lock(mutex_);
    size_t total_mem = 0;
    for (auto &mem_table_file : mem_table_file_list_) {
        total_mem += mem_table_file->GetCurrentMem();
    }
    return total_mem;
}

} // namespace engine
} // namespace milvus
} // namespace zilliz