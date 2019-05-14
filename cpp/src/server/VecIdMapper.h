/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#pragma once

#include "utils/Error.h"

#include <string>
#include <vector>
#include <unordered_map>

namespace rocksdb {
    class DB;
}

namespace zilliz {
namespace vecwise {
namespace server {

class IVecIdMapper {
public:
    static IVecIdMapper* GetInstance();

    virtual ~IVecIdMapper(){}

    virtual ServerError Put(const std::string& nid, const std::string& sid, const std::string& group = "") = 0;
    virtual ServerError Put(const std::vector<std::string>& nid, const std::vector<std::string>& sid, const std::string& group = "") = 0;

    virtual ServerError Get(const std::string& nid, std::string& sid, const std::string& group = "") const = 0;
    //NOTE: the 'sid' will be cleared at begin of the function
    virtual ServerError Get(const std::vector<std::string>& nid, std::vector<std::string>& sid, const std::string& group = "") const = 0;

    virtual ServerError Delete(const std::string& nid, const std::string& group = "") = 0;
    virtual ServerError DeleteGroup(const std::string& group) = 0;
};

class SimpleIdMapper : public IVecIdMapper{
public:
    SimpleIdMapper();
    ~SimpleIdMapper();

    ServerError Put(const std::string& nid, const std::string& sid, const std::string& group = "") override;
    ServerError Put(const std::vector<std::string>& nid, const std::vector<std::string>& sid, const std::string& group = "") override;

    ServerError Get(const std::string& nid, std::string& sid, const std::string& group = "") const override;
    ServerError Get(const std::vector<std::string>& nid, std::vector<std::string>& sid, const std::string& group = "") const override;

    ServerError Delete(const std::string& nid, const std::string& group = "") override;
    ServerError DeleteGroup(const std::string& group) override;

private:
    using ID_MAPPING = std::unordered_map<std::string, std::string>;
    mutable std::unordered_map<std::string, ID_MAPPING> id_groups_;
};

}
}
}
