const mutations = {
    // 保存token
    COMMIT_TOKEN(state, token) {
        state.token = token;
    },
    // 保存用户
    COMMIT_USER(state, user) {
        state.user = user
    },
    // 保存正在查看的题目
    SET_CUR_PROBLEM(state, problem){
        state.curProblem = problem
    },
    // 保存正在做的题目的ID
    SET_CUR_DO_PROBLEM(state, id){
        state.curDoProblemId = id
    },
    // 保存标签
    TAGS_LIST(state, arr) {
        state.tagsList = arr;
    },
    // 设置折叠
    SET_COLLAPSE(state, bool) {
        state.isCollapse = bool;
    },
    // 设置全屏
    SET_FULLSCREEN(state, bool){
        state.isFullScreen = bool
    },
    // 设置语言
    SET_LANGUAGE(state, lang) {
        state.lang = lang
    },
    // 设置面包屑
    SET_BREAD(state, breadList) {
        state.breadList = breadList
    }
}
export default mutations