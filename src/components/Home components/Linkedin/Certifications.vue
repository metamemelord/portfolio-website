<template>
  <div class="linkedin__certifications">
    <h1 style="margin: 0">Certifications</h1>
    <div class="linkedin__certifications-cards">
      <div class="linkedin__certifications-card">
        <h2>Vue JS 2 - The Complete Guide (incl. Vue Router & Vuex)</h2>
        <h4>Jan 2019 – Present</h4>
        <h3>Udemy</h3>
        <a href="https://www.udemy.com/certificate/UC-4U2B57JA/" target="blank">
          <button>
            <b>Checkout</b>
          </button>
        </a>
      </div>
      <div class="linkedin__certifications-card">
        <h2>Introduction to CSS Development</h2>
        <h4>Dec 2018 – Present</h4>
        <h3>Udemy</h3>
        <a href="https://www.udemy.com/certificate/UC-W6J6CYZ2/" target="blank">
          <button>
            <b>Checkout</b>
          </button>
        </a>
      </div>
      <div class="linkedin__certifications-card">
        <h2>GraphQL (2016)</h2>
        <h4>Nov 2018 – Present</h4>
        <h3>Udemy</h3>
        <a href="https://www.udemy.com/certificate/UC-R1QT2RZ3/" target="blank">
          <button>
            <b>Checkout</b>
          </button>
        </a>
      </div>
      <div class="linkedin__certifications-card">
        <h2>Writing High Performance Python</h2>
        <h4>Oct 2018 – Present</h4>
        <h3>Udemy</h3>
        <a href="https://www.udemy.com/certificate/UC-R3JBLJJF/" target="blank">
          <button>
            <b>Checkout</b>
          </button>
        </a>
      </div>
      <div class="linkedin__certifications-card">
        <h2>Getting started with Jenkins: Key concepts for beginners</h2>
        <h4>Aug 2018 – Present</h4>
        <h3>Udemy</h3>
        <a href="https://www.udemy.com/certificate/UC-2R70BNXL/" target="blank">
          <button>
            <b>Checkout</b>
          </button>
        </a>
      </div>
      <div class="linkedin__certifications-card">
        <h2>Deep Learning Prerequisites: The Numpy Stack in Python</h2>
        <h4>June 2018 – Present</h4>
        <h3>Udemy</h3>
        <a href="https://www.udemy.com/certificate/UC-O9Z0NBHN/" target="blank">
          <button>
            <b>Checkout</b>
          </button>
        </a>
      </div>
      <div class="linkedin__certifications-card">
        <h2>Python for Absolute Beginners</h2>
        <h4>June 2018 – Present</h4>
        <h3>Udemy</h3>
        <a href="https://www.udemy.com/certificate/UC-3RPS7V8C/" target="blank">
          <button>
            <b>Checkout</b>
          </button>
        </a>
      </div>
      <div class="linkedin__certifications-card">
        <h2>Your First Simple Android App in Java</h2>
        <h4>June 2018 – Present</h4>
        <h3>Udemy</h3>
        <a href="https://www.udemy.com/certificate/UC-PB9YQMKJ/" target="blank">
          <button>
            <b>Checkout</b>
          </button>
        </a>
      </div>
    </div>
  </div>
</template>

<style>
.linkedin__certifications {
  text-align: center;
  width: 100%;
}
.linkedin__certifications-cards {
  width: 100%;
  padding: 1rem;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-evenly;
}

.linkedin__certifications-card {
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  text-align: center;
  flex-grow: 1;
  flex-basis: 20rem;
  border-radius: 3px;
  margin: 0.5rem;
  box-shadow: 0px 0px 3px 1px var(--shadow-color);
}

.linkedin__certifications-card h2,
.linkedin__certifications-card h3 {
  margin: 0.9rem 0.7rem;
}

.linkedin__certifications-card h4 {
  margin: 0;
}

.linkedin__certifications-card button {
  bottom: 0px;
  font-size: 0.9rem;
  width: 100%;
  padding: 0.4rem;
  border: 1px transparent solid;
  color: white;
  background: rgb(255, 161, 38);
  border-radius: 0px 0px 3px 3px;
  transition: all 0.3s;
  cursor: pointer;
}

.linkedin__certifications-card button:hover {
  border: 1px solid rgb(255, 161, 38);
  background: transparent;
  color: black;
  border-radius: 0px 0px 3px 3px;
}

@media screen and (min-width: 550px) {
  .linkedin__certifications-card {
    max-width: calc(50% - 1rem);
  }
}

@media screen and (min-width: 1000px) {
  .linkedin__certifications-card {
    max-width: calc(33.33% - 1rem);
  }
}
</style>
