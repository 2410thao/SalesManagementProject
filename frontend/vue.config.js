module.exports = {
  devServer: {
    port: 3000,  // Chọn cổng 3000 cho frontend
    proxy: {
      '/api': {
        target: 'http://localhost:8080',  // Backend đang chạy trên cổng 8080
        changeOrigin: true,  // Thay đổi origin của yêu cầu (sử dụng cho các yêu cầu cross-origin)
        pathRewrite: { '^/api': '' },  // Nếu bạn muốn xóa "/api" trong đường dẫn trước khi gửi đến backend
      },
    },
  },
};
