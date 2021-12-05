// app.js
const util = require('./utils/util.js')
App({
  globalData: {
    userid: 0,
    problems: []
  },
  towxml: require('/towxml/index'),
  onLaunch() {
    this.globalData.userid = wx.getStorageSync('userid') || 0
    let that = this
    wx.getUserInfo({
      success: res => {
        util.request(wx, '/api/problem/common/userid?nick=' + res.userInfo.nickName, 'GET', {}, function (result) {
          that.globalData.userid = result.userid
          wx.setStorageSync('userid', result.userid);
        }, true)
      }
    })
    //显示转发分享
    wx.showShareMenu({
      menus: ['shareAppMessage', 'shareTimeline'],
      success(res) {},
      fail(e) {}
    })
  }
})