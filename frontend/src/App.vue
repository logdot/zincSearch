<template>
  <h1>Test</h1>
  <form v-on:submit.prevent="search">
    <input v-model="searchTerm" type="text" id="search-input" placeholder="Term to search for">
    <button>Search</button>
  </form>

  <TableComponent :fields="fields" :data="searchResults" returnField="body" @row-clicked="rowClicked"/>
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
  margin-top: 60px;
}
</style>
