<template>
  <div class="linkedin" id="linkedin">
    <div class="linkedin__label">
      <a
        class="light-a"
        href="https://www.linkedin.com/in/metamemelord/"
        target="blank"
        v-infocus="'float-animation saturated-linkedin'"
      >
        <h1>
          <i class="fa fa-linkedin"></i>
        </h1>
      </a>
    </div>
    <div class="linkedin__content">
      <linkedin-timeline/>
      <linkedin-certifications/>
      <linkedin-languages/>
      <linkedin-interests/>
    </div>
  </div>
</template>

<script>
import linkedinTimeline from "./Linkedin/Timeline.vue";
import linkedinCertifications from "./Linkedin/Certifications.vue";
import linkedinLanguages from "./Linkedin/Languages.vue";
import linkedinInterests from "./Linkedin/Interests.vue";

export default {
  components: {
    linkedinTimeline,
    linkedinCertifications,
    linkedinLanguages,
    linkedinInterests
  },
  created() {
    var dt = true,
      prev = -1;
    setInterval(() => {
      const hours = new Date().getHours();
      dt = hours > this.$store.state.day && hours < this.$store.state.night;
      if (prev != dt) {
        if (dt) {
          var label = document
            .getElementById("linkedin")
            .getElementsByTagName("a")[0];
          label.classList.remove("dark-a");
          label.classList.add("light-a");
        } else {
          var label = document
            .getElementById("linkedin")
            .getElementsByTagName("a")[0];
          label.classList.remove("light-a");
          label.classList.add("dark-a");
        }
        prev = dt;
      }
    }, 1000);
  }
};
</script>
<style>
.linkedin {
  display: flex;
  flex-wrap: wrap;
  margin: 1rem;
  border-radius: 0.3rem;
  box-shadow: 0px 0px 4px 4px var(--shadow-color);
  transition: 0.2s all;
}

.linkedin__label {
  width: 100%;
  display: flex;
  justify-content: center;
}

.linkedin__label h1 {
  margin: 0;
  font-size: 6rem;
  transition: all 5s;
  color: inherit;
}

.saturated-linkedin {
  color: #0077b5 !important;
}

.linkedin:hover {
  box-shadow: 0px 0px 3px 2px var(--shadow-color);
}

.linkedin__content {
  width: 100%;
  padding: 1rem;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}

.linkedin__content p {
  margin: 0.2rem;
  text-align: justify;
}

@media screen and (min-width: 550px) {
  .linkedin {
    max-width: 70rem;
    margin: 1rem;
  }
}

@media screen and (min-width: 73rem) {
  .linkedin {
    margin: 1rem auto;
  }
}
</style>
