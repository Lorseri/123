# get build time
MACRO(get_current_time CURRENT_TIME)
    execute_process(COMMAND "date" "+%Y-%m-%d %H:%M.%S" OUTPUT_VARIABLE ${CURRENT_TIME})
    string(REGEX REPLACE "\n" "" ${CURRENT_TIME} ${${CURRENT_TIME}})
ENDMACRO(get_current_time)

# get build type
MACRO(get_build_type)
    cmake_parse_arguments(BUILD_TYPE "" "TARGET;DEFAULT" "" ${ARGN})
    if (NOT DEFINED CMAKE_BUILD_TYPE)
        set(${BUILD_TYPE_TARGET} ${BUILD_TYPE_DEFAULT})
    elseif (CMAKE_BUILD_TYPE STREQUAL "Release")
        set(${BUILD_TYPE_TARGET} "Release")
    elseif (CMAKE_BUILD_TYPE STREQUAL "Debug")
        set(${BUILD_TYPE_TARGET} "Debug")
    else ()
        set(${BUILD_TYPE_TARGET} ${BUILD_TYPE_DEFAULT})
    endif ()
ENDMACRO(get_build_type)

# get git branch name
MACRO(get_git_branch_name GIT_BRANCH_NAME)
    set(GIT_BRANCH_NAME_REGEX "[0-9]+\\.[0-9]+\\.[0-9]")

    execute_process(COMMAND sh "-c" "git log --decorate | head -n 1 | sed 's/.*(\\(.*\\))/\\1/' | sed 's/.*, //' | sed 's=[a-zA-Z]*\/==g'"
            OUTPUT_VARIABLE ${GIT_BRANCH_NAME})

    # if (NOT GIT_BRANCH_NAME MATCHES "${GIT_BRANCH_NAME_REGEX}")
    #     execute_process(COMMAND "git" rev-parse --abbrev-ref HEAD OUTPUT_VARIABLE ${GIT_BRANCH_NAME})
    # endif ()

    if (NOT GIT_BRANCH_NAME MATCHES "${GIT_BRANCH_NAME_REGEX}")
        execute_process(COMMAND "git" symbolic-ref -q --short HEAD OUTPUT_VARIABLE ${GIT_BRANCH_NAME})
    endif ()

    message(DEBUG "GIT_BRANCH_NAME = ${GIT_BRANCH_NAME}")

    # Some unexpected case
    if (NOT GIT_BRANCH_NAME STREQUAL "")
        string(REGEX REPLACE "\n" "" GIT_BRANCH_NAME ${GIT_BRANCH_NAME})
    else ()
        set(GIT_BRANCH_NAME "#")
    endif ()
ENDMACRO(get_git_branch_name)

# get last commit id
MACRO(get_last_commit_id LAST_COMMIT_ID)
    execute_process(COMMAND sh "-c" "git log --decorate | head -n 1 | awk '{print $2}'"
            OUTPUT_VARIABLE ${LAST_COMMIT_ID})

    message(DEBUG "LAST_COMMIT_ID = ${LAST_COMMIT_ID}")

    if (NOT LAST_COMMIT_ID STREQUAL "")
        string(REGEX REPLACE "\n" "" LAST_COMMIT_ID ${LAST_COMMIT_ID})
        set(LAST_COMMIT_ID "${LAST_COMMIT_ID}")
    else ()
        set(LAST_COMMIT_ID "Unknown")
    endif ()
ENDMACRO(get_last_commit_id)

# get milvus version
MACRO(get_milvus_version)
    cmake_parse_arguments(VER "" "TARGET;DEFAULT" "" ${ARGN})

    # Step 1: get branch name
    get_git_branch_name(GIT_BRANCH_NAME)
    message(DEBUG ${GIT_BRANCH_NAME})

    # Step 2: match MAJOR.MINOR.PATCH format or set DEFAULT value
    string(REGEX MATCH "([0-9]+)\\.([0-9]+)\\.([0-9]+)" ${VER_TARGET} ${GIT_BRANCH_NAME})
    if (NOT ${VER_TARGET})
        set(${VER_TARGET} ${VER_DEFAULT})
    endif()
ENDMACRO(get_milvus_version)

# set definition
MACRO(set_milvus_definition DEF_PASS_CMAKE MILVUS_DEF)
    if (${DEF_PASS_CMAKE} STREQUAL "ON")
        add_compile_definitions(${MILVUS_DEF})
    endif()
ENDMACRO(set_milvus_definition)
