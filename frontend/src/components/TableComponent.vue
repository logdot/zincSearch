<template>
  <table id="tableComponent" class="table-fixed bg-white shadow-lg">
    <thead>
      <th v-for="field in fields" :key="field" @click="sortTable(field)" class="bg-blue-200 border text-left">
        {{field}} <i class="sort-icon" aria-label="Sort Icon"></i>
      </th>
    </thead>
    <tbody>
      <tr v-for="item in data" :key="item" @click="$emit('row-clicked', item[returnField])" class="hover:bg-gray-200">
        <td v-for="field in fields" :key="field" class="border">{{item[field]}}</td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  name: "TableComponent",
  props: {
    fields:{
      type: Array,
    },
    data:{
      type: Array,
    },
    returnField:{
      type: String,
    },
  }
}
</script>

<style scoped>

</style>