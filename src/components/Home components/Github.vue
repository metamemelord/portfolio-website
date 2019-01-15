<template>
  <div class="github">
    <div class="github__label">
      <a href="https://github.com/metamemelord" target="blank">
        <img
          src="../../assets/Octocat.png"
          alt="Github octocat"
          v-infocus="'saturate-image float-animation'"
        >
      </a>
    </div>
    <div class="github__cards">
      <div class="github__card">
        <h3>Project Name</h3>
        <p>Project description in a line</p>
        <div class="github__card-technology-label">
          <span>Java</span>
          <span>Java</span>
          <span>Java</span>
        </div>
        <button>Checkout</button>
      </div>
      <div class="github__card">
        <h3>Project Name</h3>
        <p>Project description in a line lkjdflksdjflkdsj sdksdf sdfj lksjdf sdflkj slkdjfslkdfj</p>
        <div class="github__card-technology-label">
          <span>Java</span>
          <span>Java</span>
          <span>Java</span>
        </div>
        <button>Checkout</button>
      </div>
      <div class="github__card">
        <h3>Project Name</h3>
        <p>Project description in a line</p>
        <div class="github__card-technology-label">
          <span>Java</span>
          <span>Java</span>
          <span>Java</span>
        </div>
        <button>Checkout</button>
      </div>
    </div>
  </div>
</template>

<style>
.github {
  display: flex;
  flex-wrap: wrap;
  margin: 1rem;
  border-radius: 0.3rem;
  box-shadow: 0px 0px 4px 4px rgba(0, 0, 0, 0.15);
  transition: 0.2s all;
}

.github__label {
  width: 100%;
  display: flex;
  justify-content: center;
}

.github__label img {
  margin-top: 1.4rem;
  max-width: 9rem;
  filter: grayscale(100%);
  -webkit-filter: grayscale(100%);
  -moz-filter: grayscale(100%);
  -ms-filter: grayscale(100%);
  -o-filter: grayscale(100%);
  transition: all 5s;
}

.saturate-image {
  filter: grayscale(0) !important;
  -webkit-filter: grayscale(0) !important;
  -moz-filter: grayscale(0) !important;
  -ms-filter: grayscale(0) !important;
  -o-filter: grayscale(0) !important;
}

.github:hover {
  box-shadow: 0px 0px 3px 2px rgba(0, 0, 0, 0.15);
}

.github__cards {
  width: 100%;
  padding: 1rem;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}

.github__card {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  text-align: center;
  flex-grow: 1;
  flex-basis: 20rem;
  border-radius: 3px;
  margin: 0.5rem;
  box-shadow: 0px 0px 3px 1px rgba(0, 0, 0, 0.15);
}

.github__card h3 {
  margin: 0.9rem 0.7rem;
}

.github__card p {
  margin: 0.9rem;
}

.github__card-technology-label {
  display: flex;
  justify-content: space-between;
  padding: 0 0.7rem;
  margin: 0.7rem 0;
}

.github__card-technology-label span {
  flex-grow: 1;
  text-align: center;
  padding: 0.5rem;
  margin: 0.5rem;
  background: rgb(255, 161, 38);
  border-radius: 3px;
}

.github__card button {
  bottom: 0px;
  font-size: 0.9rem;
  width: 100%;
  padding: 1rem;
  border: 1px transparent solid;
  color: white;
  background: rgb(255, 161, 38);
  border-radius: 0px 0px 3px 3px;
  transition: all 0.3s;
  cursor: pointer;
}

.github__card button:hover {
  border: 1px solid rgb(255, 161, 38);
  background: transparent;
  color: black;
  border-radius: 0px 0px 3px 3px;
}

@media screen and (min-width: 550px) {
  .github {
    max-width: 70rem;
    margin: 1rem;
  }

  .github__card {
    max-width: calc(50% - 1rem);
  }
}

@media screen and (min-width: 73rem) {
  .github {
    margin: 1rem auto;
  }
}
</style>
