<template>
  <div class="form-container">
    <h1 class="title">Thêm sản phẩm</h1>
    <form @submit.prevent="registerProduct" class="product-form">
      <div class="form-group">
        <label for="product_name">Tên sản phẩm:</label>
        <input type="text" id="product_name" v-model="product.product_name" required />
      </div>

      <div class="form-group">
        <label for="price">Giá:</label>
        <input type="number" id="price" v-model.number="product.price" required />
      </div>

      <div class="form-group">
        <label for="cost">Chi phí:</label>
        <input type="number" id="cost" v-model.number="product.cost" required />
      </div>

      <div class="form-group">
        <label for="quantity">Số lượng:</label>
        <input type="number" id="quantity" v-model.number="product.quantity" required />
      </div>

      <div class="form-group">
        <label for="unit">Đơn vị:</label>
        <input type="text" id="unit" v-model="product.unit" required />
      </div>

      <div class="form-group">
        <label for="image">Ảnh sản phẩm (URL):</label>
        <input type="text" id="image" v-model="product.image" />
      </div>

      <div class="form-group">
        <label for="category_id">Danh mục:</label>
        <input type="number" id="category_id" v-model.number="product.category_id" />
      </div>

      <div class="form-group">
        <label for="manufacturing_date">Ngày sản xuất:</label>
        <input type="date" id="manufacturing_date" v-model="product.manufacturing_date" required />
      </div>

      <div class="form-group">
        <label for="expiry_date">Ngày hết hạn:</label>
        <input type="date" id="expiry_date" v-model="product.expiry_date" required />
      </div>

      <button type="submit" class="submit-btn">Thêm sản phẩm</button>
    </form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      product: {
        product_name: "",
        price: null,
        cost: null,
        quantity: null,
        unit: "",
        image: "",
        category_id: null,
        manufacturing_date: "",
        expiry_date: "",
      },
    };
  },
  methods: {
    async registerProduct() {
      try {
        if (!this.product.category_id || isNaN(this.product.category_id)) {
          this.product.category_id = null;
        }

        const response = await axios.post("http://localhost:8080/api/products", this.product, {
          headers: { "Content-Type": "application/json" },
        });

        alert(response.data.message);
      } catch (error) {
        console.error("Lỗi khi thêm sản phẩm:", error.response?.data || error.message);
        alert("Đã xảy ra lỗi khi thêm sản phẩm.");
      }
    },
  },
};
</script>

<style scoped>
/* Container */
.form-container {
  max-width: 500px;
  margin: 30px auto;
  padding: 20px;
  background-color: #ffffff;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  border-radius: 10px;
  font-family: "Arial", sans-serif;
}

/* Tiêu đề */
.title {
  text-align: center;
  font-size: 1.8rem;
  color: #333;
  margin-bottom: 20px;
}

/* Form */
.product-form {
  display: flex;
  flex-direction: column;
}

/* Ô nhập liệu */
.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  font-weight: bold;
  margin-bottom: 5px;
  color: #555;
}

.form-group input {
  width: 100%;
  padding: 10px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 5px;
  transition: border-color 0.3s ease-in-out;
}

.form-group input:focus {
  border-color: #3498db;
  outline: none;
}

/* Nút submit */
.submit-btn {
  background-color: #3498db;
  color: white;
  padding: 12px;
  font-size: 16px;
  font-weight: bold;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease-in-out;
}

.submit-btn:hover {
  background-color: #2980b9;
}
</style>
