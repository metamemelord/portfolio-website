<template>
  <header class="page-header">
    <div class="page-header__items">
      <div class="page-header__brand">
        <span>Home</span>
      </div>
      <ul class="page-header__links">
        <li class="page-header__link">Blogs</li>
      </ul>
    </div>
  </header>
</template>

<script>
import { mapActions } from "vuex";

export default {};
</script>

<style>
.page-header {
  height: 64px;
  background: lightcoral;
}

.page-header__items {
  height: 100%;
}

.page-header__brand {
  font-size: 24px;
  padding: 0px 24px;
  display: inline-block;
  vertical-align: middle;
  font-family: "Source Sans Pro", sans-serif;
  font-weight: 700;
}
.page-header__brand span {
}

.page-header__links {
  list-style: none;
  display: inline-block;
  vertical-align: middle;
}
</style>
