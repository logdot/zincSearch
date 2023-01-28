<template>
  <div class="flex flex-col max-h-screen">
    <nav class="flex items-center justify-between flex-wrap bg-teal-500 p-6 mb-3">
      <div class="flex items-center flex-shrink-0 text-white mr-6">
        <span class="font-semibold text-xl tracking-tight">ENron VIsualizer</span>
      </div>
      <div class="w-full block flex-grow lg:flex lg:items-center lg:w-auto">
        <form v-on:submit.prevent="search" class="w-full">
          <input v-model="searchTerm" type="text" id="first_name" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="Search">
        </form>
      </div>
    </nav>

    <main class="w-full grid grid-rows-2 lg:grid-cols-2 lg:grid-rows-1 gap-6 px-3 flex-1 min-h-0 mb-3">
      <div class="overflow-y-auto overflow-x-hidden border-y">
        <TableComponent :fields="fields" :data="searchResults" returnField="body" @row-clicked="rowClicked" class="table-fixed w-full h-full"/>
      </div>

      <textarea id="body-display" v-model="bodyDisplay" class="max-h-screen w-full overflow-auto border"></textarea>
    </main>
  </div>
</template>

<script>
import axios from 'axios';
import TableComponent from "@/components/TableComponent.vue";

export default {
  name: 'App',
  components: {TableComponent},

  data() { return {
    searchTerm: '',
    searchResults: {},
    fields: ["subject", "from", "to"],
    bodyDisplay: '',
  }},

  methods: {
    search() {
      console.log(`Searching ${this.searchTerm}`)

      axios.post("http://localhost:8080/api/search", {
        search_term: this.searchTerm
      }).then((response) => {
        this.searchResults = response.data.hits.hits.map(x => x._source)
        console.log(this.searchResults)
      }).catch((error) => {
        console.error(`API error: ${error}`)
      })
    },

    rowClicked(e) {
      console.log(`Row was clicked. Returned ${e}`)

      this.bodyDisplay = e
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
