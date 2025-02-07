<template>
    <div>
      <h2>Đăng Ký Tiểu Thương</h2>
      <form @submit.prevent="registerTieuThuong">
        <div>
          <label for="tenCuaHang">Tên cửa hàng:</label>
          <input type="text" v-model="form.tenCuaHang" id="tenCuaHang" required />
        </div>
        <div>
          <label for="email">Email:</label>
          <input type="email" v-model="form.email" id="email" required />
        </div>
        <div>
          <label for="password">Mật khẩu:</label>
          <input type="password" v-model="form.password" id="password" required />
        </div>
        <div>
          <label for="soDienThoai">Số điện thoại:</label>
          <input type="text" v-model="form.soDienThoai" id="soDienThoai" />
        </div>
        <div>
          <label for="diaChi">Địa chỉ:</label>
          <input type="text" v-model="form.diaChi" id="diaChi" />
        </div>
        <div>
          <label for="tinh">Tỉnh:</label>
          <select v-model="form.tinhID" id="tinh" @change="loadCities" required>
            <option value="">Chọn tỉnh</option>
            <option v-for="tinh in tinhs" :key="tinh.id" :value="tinh.id">{{ tinh.name }}</option>
          </select>
        </div>
        <div>
          <label for="tp">Thành phố:</label>
          <select v-model="form.tpID" id="tp" @change="loadDistricts" required>
            <option value="">Chọn thành phố</option>
            <option v-for="tp in tps" :key="tp.id" :value="tp.id">{{ tp.name }}</option>
          </select>
        </div>
        <div>
          <label for="quan">Quận:</label>
          <select v-model="form.quanID" id="quan" @change="loadWards" required>
            <option value="">Chọn quận</option>
            <option v-for="quan in quans" :key="quan.id" :value="quan.id">{{ quan.name }}</option>
          </select>
        </div>
        <div>
          <label for="phuong">Phường:</label>
          <select v-model="form.phuongID" id="phuong" required>
            <option value="">Chọn phường</option>
            <option v-for="phuong in phuongs" :key="phuong.id" :value="phuong.id">{{ phuong.name }}</option>
          </select>
        </div>
        <button type="submit">Đăng ký</button>
      </form>
  
      <div v-if="message" :class="{'success': isSuccess, 'error': !isSuccess}">
        {{ message }}
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        form: {
          tenCuaHang: '',
          soDienThoai: '',
          email: '',
          password: '',
          diaChi: '',
          tinhID: '',
          tpID: '',
          quanID: '',
          phuongID: ''
        },
        tinhs: [],
        tps: [],
        quans: [],
        phuongs: [],
        message: '',
        isSuccess: false
      };
    },
    methods: {
      async loadOptions(url, fieldName) {
        try {
          const response = await axios.get(url);
          this[fieldName] = response.data;
        } catch (error) {
          console.error('Không thể tải dữ liệu:', error);
        }
      },
      loadCities() {
        this.loadOptions(`/api/locations/tps?tinh_id=${this.form.tinhID}`, 'tps');
      },
      loadDistricts() {
        this.loadOptions(`/api/locations/quans?tp_id=${this.form.tpID}`, 'quans');
      },
      loadWards() {
        this.loadOptions(`/api/locations/phuongs?quan_id=${this.form.quanID}`, 'phuongs');
      },
      async registerTieuThuong() {
        try {
          const response = await axios.post('http://localhost:8080/register-tieuthuong', {
            ten_cua_hang: this.form.tenCuaHang,
            so_dien_thoai: this.form.soDienThoai,
            email: this.form.email,
            password: this.form.password,
            dia_chi: this.form.diaChi,
            tinh_id: this.form.tinhID,
            tp_id: this.form.tpID,
            phuong_id: this.form.phuongID,
            quan_id: this.form.quanID
          });
          this.message = response.data.message;
          this.isSuccess = true;
        } catch (error) {
          this.message = error.response?.data?.error || 'Đăng ký thất bại';
          this.isSuccess = false;
        }
      }
    },
    mounted() {
      this.loadOptions('/api/locations/tinhs', 'tinhs');
    }
  };
  </script>
  
  <style scoped>
  .success {
    color: green;
  }
  .error {
    color: red;
  }
  </style>
  