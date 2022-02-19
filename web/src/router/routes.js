const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("../views/Home.vue"),
    meta: { icon: "user" },
  },
  {
    path: "/bytecalc",
    name: "ByteCalc",
    component: () => import("../views/ByteCalc.vue"),
    meta: { icon: "calculator" },
  },
];

export default routes;
