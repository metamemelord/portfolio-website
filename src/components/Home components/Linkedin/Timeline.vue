<template>
  <div class="linkedin__timeline" v-if="experiences.length != 0">
    <h1 style="margin: 0">Timeline</h1>
    <div class="linkedin__timeline-items">
      <div class="linkedin__timeline-item" v-for="(experience, idx) in experiences" :key="experience._id" :id="experience._id">
        <div class="linkedin__timeline-point" :class="{'linkedin__timeline-point-current':idx==0}"></div>
        <div class="linkedin__timeline-item-content">
          <h2>{{experience.company}}</h2>
          <h3>{{experience.title}}</h3>
          <h4>{{experience.from_date}} - {{experience.to_date}}</h4>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
        experiences: [
          {"_id":"5e0fbafc3fd85cc84aa6b7b7","company":"Zoomcar","title":"Software Engineer","from_date":"Nov, 2019","to_date":"Present"},
          {"_id":"5e0fbae03fd85cc84aa6b7b6","company":"Philips","title":"Software Engineer I","from_date":"Jul, 2018","to_date":"Nov, 2019"},
          {"_id":"5e0fbaba3fd85cc84aa6b7b5","company":"Philips","title":"Project Trainee","from_date":"Jan, 2018","to_date":"May, 2018"}
        ] 
	    }
  },
  created() {
    this.$http.get("api/experiences").then(res => {
      if (res.data.length) this.experiences = res.data
    }).catch(() => void 0);
  }
}
</script>
<style>
.linkedin__timeline {
  text-align: center;
  width: 100%;
}

.linkedin__timeline-items {
  width: 90%;
  margin-left: 10%;
}

.linkedin__timeline-point {
  position: absolute;
  left: calc(-5% - 0.9rem);
  top: calc(50% - 1rem);
  width: 2rem;
  height: 2rem;
  background: #ccc;
  border-radius: 50%;
}

.linkedin__timeline-point-current {
  background: var(--accent-color);
}

.linkedin__timeline-item {
  color: #333;
  position: relative;
  padding: 1rem;
}

.linkedin__timeline-item:first-of-type:before {
  content: "";
  position: absolute;
  top: 50%;
  height: 50%;
  left: -5%;
  width: 4px;
  background: #ccc;
}

.linkedin__timeline-item::before {
  content: "";
  position: absolute;
  top: -10%;
  height: 100%;
  left: -5%;
  width: 4px;
  background: #ccc;
}

.linkedin__timeline-item:last-of-type::before {
  content: "";
  position: absolute;
  bottom: 50%;
  height: 51%;
  left: -5%;
  width: 4px;
  background: #ccc;
}

.linkedin__timeline-item-content {
  text-align: center;
  position: relative;
  background: #ccc;
  border-radius: 0.5rem;
  padding: 1rem;
}

.linkedin__timeline-item-content h2,
.linkedin__timeline-item-content h3,
.linkedin__timeline-item-content h4 {
  margin: 0.5rem;
}

.linkedin__timeline-item-content::before {
  content: "";
  position: absolute;
  top: 45%;
  right: 100%;
  height: 0;
  width: 0;
  border: 7px solid transparent;
  border-right: 7px solid #ccc;
}
</style>
