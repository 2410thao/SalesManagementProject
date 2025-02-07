<template>
  <div class="login-container">
    <h1 class="login-title">Đăng nhập</h1>
    <form @submit.prevent="login" class="login-form">
      <input v-model="form.username" placeholder="Username" class="input-field" />
      <input v-model="form.password" type="password" placeholder="Password" class="input-field" />
      <button type="submit" class="submit-btn">Đăng nhập</button>
    </form>
    <p v-if="message" class="message">{{ message }}</p>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      form: {
        username: "",
        password: "",
      },
      message: "",
    };
  },
  methods: {
    async login() {
      try {
        const response = await axios.post("http://localhost:8080/login", this.form);
        this.message = `Đăng nhập thành công! Token: ${response.data.token}`;
        this.$emit("notify", this.message);
      } catch (error) {
        this.message = error.response?.data?.error || "Lỗi đăng nhập";
        this.$emit("notify", this.message);
      }
    },
  },
};
</script>

<style scoped>
.login-container {
  width: 100%;
  max-width: 400px;
  margin: 50px auto;
  padding: 20px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.login-title {
  text-align: center;
  font-size: 24px;
  color: #2c3e50;
  margin-bottom: 20px;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.input-field {
  padding: 12px;
  font-size: 16px;
  border: 1px solid #ddd;
  border-radius: 4px;
  outline: none;
  transition: border-color 0.3s ease;
}

.input-field:focus {
  border-color: #3498db;
}

.submit-btn {
  padding: 12px;
  font-size: 16px;
  background-color: #3498db;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.submit-btn:hover {
  background-color: #2980b9;
}

.message {
  text-align: center;
  color: #e74c3c;
  font-size: 16px;
  margin-top: 15px;
}
</style>
