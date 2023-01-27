<template>
  <table id="tableComponent">
    <thead>
      <th v-for="field in fields" :key="field" @click="sortTable(field)">
        {{field}} <i class="sort-icon" aria-label="Sort Icon"></i>
      </th>
    </thead>
    <tbody>
      <tr v-for="item in data" :key="item" @click="$emit('row-clicked', item[returnField])">
        <td v-for="field in fields" :key="field">{{item[field]}}</td>
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