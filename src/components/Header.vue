<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <router-link class="navbar-brand" to="/">Home</router-link>
    <ul class="navbar-nav mr-auto">
      <li class="nav-item">
        <router-link class="nav-link" to="/blogs">Blogs</router-link>
      </li>
    </ul>
  </nav>
</template>

<script>
import { mapActions } from "vuex";

export default {};
</script>

<style>
nav {
  height: 90;
}
</style>
