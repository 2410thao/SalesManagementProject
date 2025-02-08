<template>
  <div class="product-list-container">
    <h1 class="title">Danh sách sản phẩm</h1>
    <table class="product-table">
      <thead>
        <tr>
          <th>Id</th>
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
        <tr v-for="product in products" :key="product.product_id" @click="logProduct(product)">
          <td>{{ product.product_id }}</td>
          <td>{{ product.product_name }}</td>
          <td>{{ product.price }}</td>
          <td>{{ product.cost }}</td>
          <td>{{ product.quantity }}</td>
          <td>{{ product.unit }}</td>
          <td><img :src="product.image" alt="Product Image" class="product-image" /></td>
          <td>{{ product.category_id ?? 'Không có danh mục' }}</td>
          <td>{{ product.manufacturing_date }}</td>
          <td>{{ product.expiry_date }}</td>
          <td>
            <button @click="editProduct(product)" class="edit-btn">Sửa</button>
            <button @click="deleteProduct(product.product_id)" class="delete-btn">Xóa</button>
          </td>
        </tr>
      </tbody>
    </table>
    
    <!-- Modal chỉnh sửa sản phẩm -->
    <div v-if="editingProduct" class="modal-overlay">
      <div class="modal-container">
        <h2 class="title">Chỉnh sửa sản phẩm</h2>
        <form @submit.prevent="updateProduct" class="product-form">
          <div class="form-group">
            <label for="product_name">Tên sản phẩm:</label>
            <input type="text" id="product_name" v-model="editingProduct.product_name" required />
          </div>

          <div class="form-group">
            <label for="price">Giá:</label>
            <input type="number" id="price" v-model.number="editingProduct.price" required />
          </div>

          <div class="form-group">
            <label for="cost">Chi phí:</label>
            <input type="number" id="cost" v-model.number="editingProduct.cost" required />
          </div>

          <div class="form-group">
            <label for="quantity">Số lượng:</label>
            <input type="number" id="quantity" v-model.number="editingProduct.quantity" required />
          </div>

          <div class="form-group">
            <label for="unit">Đơn vị:</label>
            <input type="text" id="unit" v-model="editingProduct.unit" required />
          </div>

          <div class="form-group">
            <label for="image">Ảnh sản phẩm (URL):</label>
            <input type="text" id="image" v-model="editingProduct.image" />
          </div>

          <div class="form-group">
            <label for="category_id">Danh mục:</label>
            <input type="number" id="category_id" v-model.number="editingProduct.category_id" />
          </div>

          <div class="form-group">
            <label for="manufacturing_date">Ngày sản xuất:</label>
            <input type="date" id="manufacturing_date" v-model="editingProduct.manufacturing_date" required />
          </div>

          <div class="form-group">
            <label for="expiry_date">Ngày hết hạn:</label>
            <input type="date" id="expiry_date" v-model="editingProduct.expiry_date" required />
          </div>

          <button type="submit" class="submit-btn">Cập nhật</button>
          <button type="button" @click="editingProduct = null" class="cancel-btn">Hủy</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      products: [],
      editingProduct: null,
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
    editProduct(product) {
      console.log("Editing Product:", product.product_id); 
      this.editingProduct = { ...product };
      console.log(this.editingProduct.product_id); 
    },
    async updateProduct() {
      try {
        await axios.put(`http://localhost:8080/api/products/${this.editingProduct.product_id}`, this.editingProduct);
        alert("Cập nhật sản phẩm thành công!");
        this.fetchProducts();
        this.editingProduct = null;
      } catch (error) {
        console.error("Lỗi cập nhật sản phẩm:", error);
        alert("Không thể cập nhật sản phẩm.");
      }
    },
    async deleteProduct(productId) {
      if (confirm("Bạn có chắc muốn xóa sản phẩm này?")) {
        try {
          await axios.delete(`http://localhost:8080/api/products/${productId}`);
          alert("Xóa sản phẩm thành công!");
          this.fetchProducts();
        } catch (error) {
          console.error("Lỗi xóa sản phẩm:", error);
          alert("Không thể xóa sản phẩm.");
        }
      }
    },

  logProduct(product) {
    console.log("Product Data:", product);
  }
  },
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-container {
  background: #fff;
  padding: 25px;
  border-radius: 12px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
  width: 450px;
  max-width: 90%;
  text-align: center;
}

.cancel-btn {
  background-color: #e74c3c;
  color: white;
  padding: 12px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  margin-left: 10px;
}

.cancel-btn:hover {
  background-color: #c0392b;
}
</style>

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
