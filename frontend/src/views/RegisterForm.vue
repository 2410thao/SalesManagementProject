<template>
  <div>
    <h1>Đăng ký</h1>
    <form @submit.prevent="register">
      <input v-model="form.username" placeholder="Username" />
      <input v-model="form.password" type="password" placeholder="Password" />
      <button type="submit">Đăng ký</button>
    </form>
    <p v-if="message">{{ message }}</p>
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
    async register() {
      try {
        const response = await axios.post("http://localhost:8080/register", this.form);
        this.message = response.data.message;
        this.$emit("notify", response.data.message);
      } catch (error) {
        this.message = error.response?.data?.error || "Lỗi đăng ký";
        this.$emit("notify", this.message);
      }
    },
  },
};
</script>
<style scoped>
div {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background-color: #f4f4f4;
}

h1 {
  font-size: 24px;
  margin-bottom: 20px;
  color: #333;
}

form {
  display: flex;
  flex-direction: column;
  gap: 10px;
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  width: 300px;
}

input {
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 16px;
}

button {
  padding: 10px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
  transition: background 0.3s;
}

button:hover {
  background: #0056b3;
}

p {
  margin-top: 10px;
  font-size: 14px;
  color: green;
}
</style>
