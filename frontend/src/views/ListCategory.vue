<template>
    <div class="category-list">
      <h1>Danh sách danh mục</h1>
      <button @click="goToAddCategory" class="add-btn">Thêm danh mục</button>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Tên danh mục</th>
            <th>Mô tả</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="category in categories" :key="category.category_id">
            <td>{{ category.category_id }}</td>
            <td>{{ category.name }}</td>
            <td>{{ category.description }}</td>
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
        categories: []
      };
    },
    methods: {
      async fetchCategories() {
        try {
          const response = await axios.get("http://localhost:8080/api/categories");
          this.categories = response.data.categories;
        } catch (error) {
          console.error("Lỗi khi lấy danh sách danh mục:", error);
        }
      },
      goToAddCategory() {
        this.$router.push("/addcategory");
      }
    },
    mounted() {
      this.fetchCategories();
    }
  };
  </script>
  
  <style scoped>
  .category-list {
    max-width: 800px;
    margin: 20px auto;
  }
  .add-btn {
    background-color: #007bff;
    color: white;
    padding: 8px;
    border: none;
    cursor: pointer;
    margin-bottom: 10px;
  }
  table {
    width: 100%;
    border-collapse: collapse;
  }
  th, td {
    border: 1px solid #ddd;
    padding: 8px;
  }
  th {
    background-color: #f4f4f4;
  }
  </style>
  