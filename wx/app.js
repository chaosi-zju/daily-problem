// app.js
App({
  globalData: {
    problems: []
  },
  towxml: require('/towxml/index'),
  onLaunch() {
    // 登录
    wx.login({
      success: res => {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      }
    })
    //显示转发分享
    wx.showShareMenu()
  },
  shareAppMessage: function() {
    //转发
    return {
      title: '快来一起刷题吧，每天三道算法题，面试不是问题！',
      path: 'pages/index/index',
      imageUrl: '/image/fan.jpg',
      success: function(res) {
        util.toast(wx, '转发成功', 800)
      },
      fail: function(res) {
        util.toast(wx, '转发失败', 800)
      }
    }
  },
})