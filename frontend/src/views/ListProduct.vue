<template>
  <div class="product-list-container">
    <h1 class="title">Danh sách sản phẩm</h1>
    <table class="product-table">
      <thead>
        <tr>
          <th>Tên sản phẩm</th>
          <th>Giá</th>
          <th>Chi phí</th>
          <th>Số lượng</th>
          <th>Đơn vị</th>
          <th>Ảnh</th>
          <th>Danh mục</th>
          <th>Ngày sản xuất</th>
          <th>Ngày hết hạn</th>
          <th>Thao tác</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="product in products" :key="product.product_name">
          <td>{{ product.product_name }}</td>
          <td>{{ product.price }}</td>
          <td>{{ product.cost }}</td>
          <td>{{ product.quantity }}</td>
          <td>{{ product.unit }}</td>
          <td><img :src="product.image" alt="Product Image" class="product-image" /></td>
          <td>{{ product.category_id ?? 'Không có danh mục' }}</td>
          <td>{{ product.manufacturing_date }}</td>
          <td>{{ product.expiry_date }}</td>
          <td></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      products: [],
    };
  },
  mounted() {
    this.fetchProducts();
  },
  methods: {
    async fetchProducts() {
      try {
        const response = await axios.get("http://localhost:8080/api/products");
        this.products = response.data.products;
      } catch (error) {
        console.error("Lỗi tải danh sách sản phẩm:", error);
        alert("Không thể tải danh sách sản phẩm.");
      }
    },
  },
};
</script>

<style scoped>
.product-list-container {
  font-family: 'Arial', sans-serif;
  margin: 20px;
}

.title {
  text-align: center;
  font-size: 2rem;
  color: #2c3e50;
  margin-bottom: 20px;
}

.product-table {
  width: 100%;
  border-collapse: collapse;
  margin: 0 auto;
  background-color: #fff;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.product-table th,
.product-table td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.product-table th {
  background-color: #3498db;
  color: white;
  font-weight: bold;
}

.product-table td {
  font-size: 14px;
  color: #7f8c8d;
}

.product-image {
  width: 50px;
  height: 50px;
  object-fit: cover;
  border-radius: 5px;
}

.product-table tr:nth-child(even) {
  background-color: #f9f9f9;
}
</style>
