import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import Login from '@/views/LoginForm.vue'; // Đường dẫn đến Login.vue
import Register from '@/views/RegisterForm.vue'; // Đường dẫn đến Register.vue
import Dktieuthuong from '@/views/RegisterTieuthuong.vue'; // Đường dẫn đến RegisterTieuthuong.vue
import Addsanpham from '@/views/AddProduct.vue'; // Đường dẫn đến RegisterTieuthuong.vue
import Listproduct from '@/views/ListProduct.vue'; 
import AddCategory from '@/views/AddCategory.vue'; // Import trang thêm danh mục
import ListCategory from '@/views/ListCategory.vue'; // Import trang danh mục
const routes: Array<RouteRecordRaw> = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
  },
  {
    path: '/dktieuthuong',
    name: 'Dktieuthuong',
    component: Dktieuthuong,  // Đổi từ Register thành Dktieuthuong
  },
  {
    path: '/addsanpham',
    name: 'Addsanpham',
    component: Addsanpham,  // Đổi từ Register thành Dktieuthuong
  },
  {
    path: '/listproduct',
    name: 'Listproduct',
    component: Listproduct,
  },
  {
    path: '/addcategory',
    name: 'AddCategory',
    component: AddCategory,
  }, 
  {
    path: '/listcategory',
    name: 'ListCategory',
    component: ListCategory,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
