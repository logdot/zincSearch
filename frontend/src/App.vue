<template>
  <h1>Test</h1>
  <form v-on:submit.prevent="search">
    <input v-model="searchTerm" type="text" id="search-input" placeholder="Term to search for">
    <button>Search</button>
  </form>
</template>

<script>
import axios from 'axios';

export default {
  name: 'App',

  data() { return {
    searchTerm: '',
    searchResults: '',
  }},

  methods: {
    search() {
      console.log(`Searching ${this.searchTerm}`)

      axios.post("http://localhost:8080/api/search", {
        search_term: this.searchTerm
      }).then((response) => {
        this.searchResults = response.data.results
        console.log(this.searchResults)
      }).catch((error) => {
        console.error(`API error: ${error}`)
      })
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
