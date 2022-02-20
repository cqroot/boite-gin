<template>
  <div id="page-bytecalc">
    <a-row
      v-for="item in byteData"
      :key="item.name"
      style="margin-bottom: 10px; height: 40px"
    >
      <a-col :span="8">
        <div style="text-align: right; padding-right: 10px; line-height: 40px">
          {{ item.text }}:
        </div>
      </a-col>
      <a-col :span="12">
        <a-input
          v-model="item.value"
          style="height: 40px"
          @change="inputChange(item.name, item.value)"
          @pressEnter="calculate(item.name, item.value)"
        />
      </a-col>
    </a-row>

    <a-row>
      <a-col :span="8"> </a-col>
      <a-col :span="12">
        <a-button
          type="primary"
          style="margin-right: 10px"
          @click="calculate(currentUnit, currentValue)"
        >
          Calculate
        </a-button>
        <a-button @click="clear"> Clear </a-button>
      </a-col>
    </a-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currentUnit: "",
      currentValue: "",
      byteData: [
        {
          name: "bit",
          text: "bit (b)",
          value: "",
        },
        {
          name: "byte",
          text: "byte (B)",
          value: "",
        },
        {
          name: "kilobyte",
          text: "kilobyte (kB)",
          value: "",
        },
        {
          name: "megabyte",
          text: "megabyte (MB)",
          value: "",
        },
        {
          name: "gigabyte",
          text: "gigabyte (GB)",
          value: "",
        },
        {
          name: "terabyte",
          text: "terabyte (TB)",
          value: "",
        },
      ],
    };
  },
  methods: {
    inputChange(name, value) {
      this.currentUnit = name;
      this.currentValue = value;
    },
    calculate(name, value) {
      this.axios
        .get("/api/bytecalc", {
          params: {
            unit: name,
            value: value,
          },
        })
        .then((response) => {
          if (response.data.code != 0) {
            this.$message.info(response.data);
          } else {
            this.byteData.forEach((item) => {
              item.value = response.data.data[item.name];
            });
          }
        });
    },
    clear() {
      this.byteData.forEach((item) => (item.value = ""));
    },
  },
};
</script>

<style lang="scss" scoped>
#page-bytecalc {
  width: 100%;
  min-width: 350px;
  position: absolute;
  left: 50%;
  top: 50%;
  -webkit-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
}
</style>
