import {
    createApp,
    reactive
} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import Loading from './components/Loading.vue'
import Navbar from './components/Navbar.vue'
import Photo from './components/Photo.vue'
import UserMiniCard from './components/UserMiniCard.vue'
import PageNotFound from './components/404.vue'
import Like from './components/Like.vue'
import Comment from './components/Comment.vue'
import PhotoComment from './components/PhotoComment.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;

app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", Loading);
app.component("Navbar", Navbar);
app.component("Photo", Photo);
app.component("UserMiniCard", UserMiniCard);
app.component("PageNotFound", PageNotFound);
app.component("LikeModal", Like);
app.component("CommentModal", Comment);
app.component("PhotoComment", PhotoComment);
app.use(router)
app.mount('#app')