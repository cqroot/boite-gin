const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("../views/Home.vue"),
    meta: { icon: "user" },
  },
  {
    path: "/byte_calc",
    name: "ByteCalc",
    component: () => import("../views/ByteCalc.vue"),
    meta: { icon: "user" },
  },
];

export default routes;
