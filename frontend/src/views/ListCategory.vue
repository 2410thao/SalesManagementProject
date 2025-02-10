<template>
  <div class="category-list-container">
    <h1 class="title">Danh sách danh mục</h1>
    <table class="category-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Tên danh mục</th>
          <th>Mô tả</th>
          <th>Thao tác</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="category in categories" :key="category.category_id">
          <td>{{ category.category_id }}</td>
          <td>{{ category.name }}</td>
          <td>{{ category.description }}</td>
          <td>
            <button @click="editCategory(category)" class="edit-btn">Sửa</button>
            <button @click="deleteCategory(category.category_id)" class="delete-btn">Xóa</button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Modal chỉnh sửa danh mục -->
    <div v-if="editingCategory" class="modal-overlay">
      <div class="modal-container">
        <h2 class="title">Chỉnh sửa danh mục</h2>
        <form @submit.prevent="updateCategory" class="category-form">
          <div class="form-group">
            <label for="name">Tên danh mục:</label>
            <input type="text" id="name" v-model="editingCategory.name" required />
          </div>

          <div class="form-group">
            <label for="description">Mô tả:</label>
            <textarea id="description" v-model="editingCategory.description" required></textarea>
          </div>

          <button type="submit" class="submit-btn">Cập nhật</button>
          <button type="button" @click="editingCategory = null" class="cancel-btn">Hủy</button>
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
      categories: [],
      editingCategory: null,
    };
  },
  mounted() {
    this.fetchCategories();
  },
  methods: {
    // Lấy danh sách danh mục từ API
    async fetchCategories() {
      try {
        const response = await axios.get("http://localhost:8080/api/categories");
        this.categories = response.data.categories;
      } catch (error) {
        console.error("Lỗi khi tải danh mục:", error);
        alert("Không thể tải danh sách danh mục.");
      }
    },

    // Mở modal chỉnh sửa danh mục
    editCategory(category) {
      this.editingCategory = { ...category };
    },

    // Cập nhật danh mục
    async updateCategory() {
      try {
        await axios.put(`http://localhost:8080/api/categories/${this.editingCategory.category_id}`, this.editingCategory);
        alert("Cập nhật danh mục thành công!");
        this.fetchCategories();
        this.editingCategory = null;
      } catch (error) {
        console.error("Lỗi cập nhật danh mục:", error);
        alert("Không thể cập nhật danh mục.");
      }
    },

    // Xóa danh mục
    async deleteCategory(categoryId) {
      if (confirm("Bạn có chắc muốn xóa danh mục này?")) {
        try {
          await axios.delete(`http://localhost:8080/api/categories/${categoryId}`);
          alert("Xóa danh mục thành công!");
          this.fetchCategories();
        } catch (error) {
          console.error("Lỗi xóa danh mục:", error);
          alert("Không thể xóa danh mục.");
        }
      }
    },
  },
};
</script>
<style scoped>
.category-list-container {
  font-family: 'Arial', sans-serif;
  margin: 20px;
}

.title {
  text-align: center;
  font-size: 2rem;
  color: #2c3e50;
  margin-bottom: 20px;
}

.category-table {
  width: 100%;
  border-collapse: collapse;
  margin: 0 auto;
  background-color: #fff;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.category-table th,
.category-table td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.category-table th {
  background-color: #3498db;
  color: white;
  font-weight: bold;
}

.category-table td {
  font-size: 14px;
  color: #7f8c8d;
}

.category-table tr:nth-child(even) {
  background-color: #f9f9f9;
}

/* Modal CSS */
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
  width: 400px;
  max-width: 90%;
  text-align: center;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  font-weight: bold;
  margin-bottom: 5px;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.submit-btn {
  background-color: #27ae60;
  color: white;
  padding: 12px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
}

.submit-btn:hover {
  background-color: #219150;
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
