#pragma once

#include <src/pb/message.pb.h>
#include "dog_segment/Partition.h"
#include "SegmentDefs.h"

namespace milvus::dog_segment {

class Collection {
public:
    explicit Collection(std::string &collection_name, std::string &schema);

    void AddIndex(const grpc::IndexParam &index_param);

    void CreateIndex(std::string &index_config);

    void parse();

public:
    SchemaPtr& get_schema() {
      return schema_;
    }

    IndexMetaPtr& get_index() {
      return index_;
    }

    std::string& get_collection_name() {
      return collection_name_;
    }

private:
    IndexMetaPtr index_;
    std::string collection_name_;
    std::string schema_json_;
    SchemaPtr schema_;
};

using CollectionPtr = std::unique_ptr<Collection>;

}