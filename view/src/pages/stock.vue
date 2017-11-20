<template>
  <div class="stockPage">
    <h1>Stock Page should have these component</h1>
    <div class="container">
      <md-field>
        <label>Search Stock</label>
        <md-input v-model="searchStock"></md-input>
      </md-field>

      <md-card>
        <md-card-header>
          <md-card-header-text>
            <div class="md-title">Stock Title</div>
            <div class="md-subhead">Stock Detail</div>
          </md-card-header-text>

          <md-card-media>
            <img src="https://image.flaticon.com/icons/svg/25/25345.svg" alt="Stock">
          </md-card-media>
        </md-card-header>

        <md-card-actions>
          <md-button>Add To List</md-button>
          <md-button>Remove</md-button>
        </md-card-actions>
      </md-card>
    </div>

  </div>
</template>

<script>
  export default {
    name: 'stock',
    data () {
      return {
        searchStock: '',
      }
    }
  }
</script>

<style lang="scss" scoped>
  .stockPage {
    .container {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
    }

    .md-field {
      max-width: 300px;
    }
  }
</style>
