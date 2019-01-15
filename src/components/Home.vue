<template>
  <main>
    <div class="main__intro">
      <div class="main__intro-image-container">
        <img class="main__intro-image" src="../assets/DP.png" alt="Gaurav Saini">
      </div>
      <div class="main__intro-content">
        <h1>Gaurav Saini</h1>
        <p>I learn things by doing.</p>
        <div class="main__intro-information">
          <span>
            <i class="fa fa-mobile"></i>
            <span>+91-99994-97257</span>
          </span>
          <span>
            <i class="fa fa-location-arrow"></i>
            <span>Bangalore</span>
          </span>
        </div>
      </div>
    </div>
    <div class="github">
      <div class="github__label">
        <h1>
          <i class="fa fa-github"></i>
        </h1>
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
  </main>
</template>

<script>
export default {};
</script>

<style>
main {
  max-width: 70rem;
}
.main__intro {
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: space-between;
  margin: 1rem;
  border-radius: 0.3rem;
  box-shadow: 0px 0px 4px 4px rgba(0, 0, 0, 0.15);
  transition: 0.2s all;
}

.main__intro:hover {
  box-shadow: 0px 0px 3px 2px rgba(0, 0, 0, 0.15);
}

.main__intro-image {
  margin: 2rem;
  margin-bottom: 1rem;
  height: 40vmin;
  max-height: 20rem;
  min-height: 10rem;
}

.main__intro-content {
  text-align: center;
  padding: 1rem;
  padding-top: 0;
  flex-grow: 1;
}

.main__intro-content > h1 {
  margin-top: 0;
  font-size: 3rem;
}

.main__intro-information {
  display: flex;
  justify-content: center;
}

.main__intro-information > span * {
  margin: 0.3rem;
  vertical-align: middle;
}

.main__intro-information i {
  color: rgb(255, 161, 38);
  font-size: 2rem;
}

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

.github__label h1 {
  margin: 0;
  font-size: 7rem;
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
  text-align: center;
  flex-grow: 1;
  flex-basis: 20rem;
  /* padding: 0.7rem; */
  border-radius: 3px;
  margin: 0.5rem;
  box-shadow: 0px 0px 3px 1px rgba(0, 0, 0, 0.15);
}

.github__card h3 {
  margin: 0.9rem 0.7rem;
}

.github__card p {
  margin: 0.9rem 0;
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
  main {
    margin: 0 auto;
  }
  .main__intro {
    flex-direction: row;
    max-width: 70rem;
    margin: 1rem;
  }
  .main__intro-image {
    margin: 2rem;
    height: 25vmin;
  }
  .main__intro-content h1 {
    margin-top: inherit;
  }

  .github {
    max-width: 70rem;
    margin: 1rem;
  }
}

@media screen and (min-width: 70rem) {
  .main__intro {
    margin: 1rem auto;
  }
  .github {
    margin: 1rem auto;
  }
}
</style>

