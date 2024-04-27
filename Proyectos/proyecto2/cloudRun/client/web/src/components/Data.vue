<template>
    <div>
      <h1>Data from Mongo</h1>
      <button @click="fetchData">Load More</button>
      <div v-if="items.length > 0" class="scrollable-list">
        <ul>
          <!--map-->
          <li v-for="item in items" :key="item._id">
            <p>{{ item.message}}</p>
          </li>
        </ul>
    </div>
    <div v-else>
      No data available.
    </div>
      
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        items: []
      };
    },
    methods: {
      async fetchData() {
        try {
            const response = await axios.get(`https://api-cl7ewgkprq-uc.a.run.app/api/votes`);
            //save into temp array
            const temp = response.data;
            //get the last 20 items
            const lastItems = temp.slice(-20);
            //save the 20 items to the items array
            this.items = lastItems;
        } catch (error) {
          console.error('Error fetching data:', error);
        }
      }
    },
    mounted() {
      this.fetchData();
    }
  };
  </script>
  
  <style scoped>
  .scrollable-list {
  max-height: 400px; /* Adjust the max-height as needed */
  overflow-y: auto;
}

.scrollable-list ul {
    list-style-type: none;
    padding: 0;
}

.scrollable-list li {
    padding: 10px;
    border-bottom: 1px solid #ccc;

}

.scrollable-list li:last-child {
    border-bottom: none;
}
  </style>