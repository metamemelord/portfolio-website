<template>
  <div class="linkedin__timeline">
    <h1 style="margin: 0">Timeline</h1>
    <div class="linkedin__timeline-items">
      <div class="linkedin__timeline-item">
        <div class="linkedin__timeline-point linkedin__timeline-point-current"></div>
        <div class="linkedin__timeline-item-content">
          <h2>Philips India</h2>
          <h3>Software Engineer 1</h3>
          <h4>July, 2018 - present</h4>
        </div>
      </div>
      <div class="linkedin__timeline-item">
        <div class="linkedin__timeline-point"></div>
        <div class="linkedin__timeline-item-content">
          <h2>Philips India</h2>
          <h3>Project Trainee</h3>
          <h4>Jan, 2018 - May, 2018</h4>
        </div>
      </div>
    </div>
  </div>
</template>

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
  background: rgb(255, 161, 38);
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
  height: 50%;
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
