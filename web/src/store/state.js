const state = {
    token: '',              //用户token
    user: '',               //用户信息
    curProblem: {},         //当前查看的题目
    curDoProblemId: '',     //当前正在做的题目的ID
    lang: 'zh',             //默认语言
    isCollapse: false,      //侧边导航是否折叠
    isFullScreen: false,    //是否全屏
    tagsList: [],           //打开的标签页个数,
    breadList: ['home']     //面包屑导航
}
export default state