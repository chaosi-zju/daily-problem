// const host = 'https://sslapi.chaosi-zju.com'
const host = 'http://127.0.0.1:5001'

const request = function (wx, path, method, data, func, noloading) {
  if (noloading == null || !noloading) {
    loading(wx, 'loading...')
  }
  wx.request({
    url: host + path,
    data: data,
    method: method,
    header: {
      'content-type': 'application/json'
    },
    success: function (res) {
      if (noloading == null || !noloading) {
        wx.hideLoading()
      }
      if (res.data.code == 200) {
        // 从服务器获取到相应数据
        var result = res.data.data;
        func(result);
      } else {
        if(res.data.message !== '') {
          toast(wx, res.data.message, 800)
        }else{
          toast(wx, '服务器有问题，数据返回有误', 800)
        }
      }
    },
    fail: function () {
      if (noloading == null || !noloading) {
        wx.hideLoading()
      }
      toast(wx, '网络错误，请稍后再试', 1000)
    }
  })
}

const loading = function (wx, title) {
  wx.showLoading({
    title: title,
    mask: true
  })
}

const toast = function (wx, title, time) {
  wx.showToast({
    title: title,
    icon: 'none',
    duration: time
  })
}

module.exports = {
  request: request,
  loading: loading,
  toast: toast,
  host: host,
}