<template>
  <div class="linkedin__languages">
    <h1 style="margin: 0">Languages</h1>
    <div class="linkedin__languages-items">
      <span>English</span>
      <span>Hindi</span>
      <span>Spanish</span>
    </div>
  </div>
</template>

<style>
.linkedin__languages {
  text-align: center;
  width: 100%;
}

.linkedin__languages-items {
  display: flex;
  justify-content: center;
  margin: 1rem;
}

.linkedin__languages-items span {
  margin: 0.5rem;
  padding: 0.5rem;
  background: var(--accent-color);
  color: #2e3342;
  border-radius: 0.2rem;
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>
