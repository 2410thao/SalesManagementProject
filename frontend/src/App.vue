<!-- <template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App',
};
</script>

<style>
/* Bạn có thể thêm CSS nếu cần */
</style> -->
<!-- <template>
  <div>
    <h1>Đăng ký</h1>
    <form @submit.prevent="register">
      <input v-model="registerForm.username" placeholder="Username" />
      <input v-model="registerForm.password" type="password" placeholder="Password" />
      <button type="submit">Đăng ký</button>
    </form>

    <h1>Đăng nhập</h1>
    <form @submit.prevent="login">
      <input v-model="loginForm.username" placeholder="Username" />
      <input v-model="loginForm.password" type="password" placeholder="Password" />
      <button type="submit">Đăng nhập</button>
    </form>

    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      registerForm: {
        username: "",
        password: "",
      },
      loginForm: {
        username: "",
        password: "",
      },
      message: "",
    };
  },
  methods: {
    async register() {
      try {
        const response = await axios.post("http://localhost:8080/register", this.registerForm);
        this.message = response.data.message;
      } catch (error) {
        this.message = error.response.data.error || "Lỗi đăng ký";
      }
    },
    async login() {
      try {
        const response = await axios.post("http://localhost:8080/login", this.loginForm);
        this.message = `Đăng nhập thành công! Token: ${response.data.token}`;
      } catch (error) {
        this.message = error.response.data.error || "Lỗi đăng nhập";
      }
    },
  },
};
</script> -->
<!-- <template>
  <div>
    <nav>
      <router-link to="/login">Đăng nhập</router-link>
      <router-link to="/register">Đăng ký</router-link>
    </nav>
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
nav {
  display: flex;
  gap: 20px;
  margin-bottom: 20px;
}
a {
  text-decoration: none;
  color: blue;
}
a.router-link-active {
  font-weight: bold;
  color: darkblue;
}
</style> -->
<template>
  <div id="app" class="app-container">
    <header class="header">
      <h1>Welcome to Our Shop</h1>
      <nav class="nav">
        <router-link to="/login" class="nav-link">Đăng nhập</router-link>
        <router-link to="/register" class="nav-link">Đăng ký</router-link>
        <router-link to="/addsanpham" class="nav-link">Thêm sản phẩm</router-link>
        <router-link to="/listproduct" class="nav-link">Danh sách sản phẩm</router-link>
        <router-link to="/addcategory" class="nav-link">Thêm danh mục</router-link>
        <router-link to="/listcategory" class="nav-link">Danh sách danh mục</router-link>
      </nav>
    </header>

    <main class="main-content">
      <router-view></router-view>
    </main>

    <footer class="footer">
      <p>&copy; 2025 Shop Online. All rights reserved.</p>
    </footer>
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  font-family: Arial, sans-serif;
  background-color: #f9f9f9;
  color: #333;
}

.header {
  background-color: #4CAF50;
  color: white;
  padding: 1rem;
  text-align: center;
}

.nav {
  display: flex;
  justify-content: center;
  gap: 1rem;
  margin-top: 0.5rem;
}

.nav-link {
  text-decoration: none;
  color: white;
  font-weight: bold;
}

.nav-link:hover {
  text-decoration: underline;
}

.main-content {
  flex: 1;
  padding: 2rem;
  text-align: center;
}

.footer {
  background-color: #333;
  color: white;
  text-align: center;
  padding: 1rem;
}
</style>
