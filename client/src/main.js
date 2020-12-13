import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import firebase from "firebase/app";
Vue.config.productionTip = false;

// ここはオープンな情報らしい？？
const firebaseConfig = {
  apiKey: "AIzaSyC4W93WffJfa3Mpc-mPKRODVrB_fkkxNHU",
  authDomain: "testproj-eb320.firebaseapp.com",
  projectId: "testproj-eb320",
  storageBucket: "testproj-eb320.appspot.com",
  messagingSenderId: "1089861809595",
  appId: "1:1089861809595:web:c0e3776ccbd6cec671c09e",
  measurementId: "G-TEKXDYGDTR"
};

firebase.initializeApp(firebaseConfig);

let app;

firebase.auth().onAuthStateChanged(() => {
  /* eslint-disable no-new */
  if (!app) {
    new Vue({
      router,
      render: h => h(App)
    }).$mount("#app");
  }
});
