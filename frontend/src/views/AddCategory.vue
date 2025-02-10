<template>
    <div class="add-category-container">
      <h2>Thêm mới danh mục</h2>
      <form @submit.prevent="addCategory">
        <div class="form-group">
          <label for="name">Tên danh mục:</label>
          <input type="text" id="name" v-model="category.name" required />
        </div>
  
        <div class="form-group">
          <label for="description">Mô tả:</label>
          <textarea id="description" v-model="category.description"></textarea>
        </div>
  
        <button type="submit" class="submit-btn">Thêm danh mục</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  
  export default {
    data() {
      return {
        category: {
          name: "",
          description: ""
        }
      };
    },
    methods: {
      async addCategory() {
        try {
          const response = await axios.post("http://localhost:8080/api/categories", this.category);
          alert("Thêm danh mục thành công!");
          this.category.name = "";
          this.category.description = "";
        } catch (error) {
          console.error("Lỗi khi thêm danh mục:", error);
          alert("Không thể thêm danh mục.");
        }
      }
    }
  };
  </script>
  
  <style scoped>
  .add-category-container {
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    background: #fff;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  }
  
  .form-group {
    margin-bottom: 15px;
  }
  
  label {
    display: block;
    font-weight: bold;
    margin-bottom: 5px;
  }
  
  input, textarea {
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  
  .submit-btn {
    background-color: #28a745;
    color: white;
    padding: 10px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  
  .submit-btn:hover {
    background-color: #218838;
  }
  </style>
  