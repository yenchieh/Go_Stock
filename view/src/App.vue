<template>
  <div id="app" class="page-container">
    <md-app>
      <md-app-toolbar class="md-primary">
        <span class="md-title">
          <img class="logo" src="/assets/images/logo_white.svg" height="30px" />
          Go Stock
        </span>
      </md-app-toolbar>

      <md-app-content class="container">
        <router-view></router-view>
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
export default {
  name: 'app',
  data () {
    return {
    }
  }
}
</script>

<style lang="scss" scoped>
  .container {
    text-align: center;
  }
</style>

<style lang="scss">
  @import "./styles/colors";
  @import "~vue-material/dist/theme/engine"; // Import the theme engine

  @include md-register-theme("default", (
    primary: $primary, // The primary color of your brand
    accent: $accent // The secondary color of your brand
  ));

  @import "~vue-material/dist/theme/all"; // Apply the theme

  #app {
    img.logo {
      height: 40px;
    }
  }
</style>
