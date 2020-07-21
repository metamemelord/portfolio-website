<template>
  <transition name="grow" mode="out-in">
    <div v-if="contactMeDialog" @click.self="closeContactMeDialog()" class="contact-me-dialog__wrapper">
      <div v-if="contactMeDialog" class="contact-me-dialog">
        <div>
          <div class="contact-me-dialog-closer" @click="closeContactMeDialog()">
            <i class="fa fa-times" aria-hidden="true"></i>
          </div>
        </div>
        <template v-if="form_status==0">
          <h1>Say hello!</h1>
          <form>
            <input
              type="text"
              @focus="invalid_name=false"
              :class="{'invalid-content': invalid_name}"
              v-model="name"
              placeholder="Full name"
            />
            <input
              type="email"
              @focus="invalid_email=false"
              :class="{'invalid-content': invalid_email}"
              v-model="email"
              placeholder="Email"
            />
            <textarea
              @focus="invalid_body=false"
              :class="{'invalid-content': invalid_body}"
              v-model="body"
              class="contact-me-dialog__body"
              row="7"
              placeholder="What is this about? (Max 500 characters)"
            />
            <input
              :disabled="send_queued"
              class="contact-me-dialog__submit-btn"
              type="submit"
              @click.prevent="submitForm()"
              value="Submit"
            />
          </form>
        </template>
        <template v-else-if="form_status==1">
          <h1 class="contact-me-dialog__status">
            <i style="color:red;" class="fa fa-times-circle-o"></i>
          </h1>
          <span style="margin-bottom: 2rem;">Something's not right, please try again later.</span>
        </template>
        <template v-else>
          <h1 class="contact-me-dialog__status">
            <i class="fa fa-check-circle-o"></i>
          </h1>
          <span style="margin-bottom: 2rem;">Your message has been sent</span>
        </template>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: "contact-me",
  data() {
    return {
      name: "",
      email: "",
      body: "",
      invalid_name: false,
      invalid_email: false,
      invalid_body: false,
      send_queued: false,
      form_status: 0
    };
  },
  computed: {
    contactMeDialog() {
      return this.$store.state.contactMeDialog;
    }
  },
  methods: {
    closeContactMeDialog() {
      this.form_status = 0;
      this.name = "";
      this.email = "";
      this.body = "";
      this.send_queued = false;
      this.form_status = 0;
      this.$store.dispatch("setContactMeDialog", false);
    },
    submitForm() {
      const formData = new FormData();

      if (!this.name) {
        this.invalid_name = true;
        return;
      }
      formData.append("name", this.name);

      if (!this.email) {
        this.invalid_email = true;
        return;
      } else if (
        !/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(
          this.email
        )
      ) {
        this.invalid_email = true;
        return;
      }
      formData.append("email", this.email);

      if (!this.body) {
        this.invalid_body = true;
        return;
      } else if (this.body.length > 500) {
        this.invalid_body = true;
        this.body = this.body.substring(0, 500);
        return;
      }
      formData.append("datetime", new Date());
      formData.append("body", this.body);

      this.send_queued = true;
      this.$http
        .post("/api/email", formData)
        .then(() => {
          this.form_status = 2;
        })
        .catch(() => {
          this.form_status = 1;
        });
    }
  }
};
</script>

<style scoped>
.invalid-content,
.contact-me-dialog input:invalid {
  background: rgba(255, 91, 91, 0.15);
  border: 1px solid red !important;
}

.contact-me-dialog__wrapper {
  position: fixed;
  display: flex;
  top: 0;
  background: transparent;
  justify-content: center;
  height: 100vh;
  width: 100%;
  padding-top: 6.3rem;
  z-index: 2;
}

.contact-me-dialog {
  position: relative;
  width: 100%;
  max-height: 28.5rem;
  padding: 0.5rem;
  margin: 1rem;
  background: var(--background-color);
  color: var(--text-color);
  display: flex;
  align-items: center;
  flex-flow: column;
  border-radius: 0.5rem;
  box-shadow: 0px 0px 12px 12px var(--shadow-color);
  z-index: 15;
  max-width: 70rem;
  min-width: 23rem;
}

.contact-me-dialog form {
  text-align: center;
  width: 100%;
  padding: 0.5rem;
}

.contact-me-dialog input {
  margin: 0.5rem;
  width: 90%;
  padding: 0.5rem;
  border-radius: 0.2rem;
  color: var(--text-color);
  border: 1px solid var(--text-color);
  background: var(--background-color);
}

.contact-me-dialog input:focus,
.contact-me-dialog__body:focus {
  border: 1px solid var(--accent-color);
  border-radius: 0.2rem;
}

.contact-me-dialog input::placeholder,
.contact-me-dialog input::-ms-input-placeholder,
.contact-me-dialog__body::placeholder,
.contact-me-dialog__body::-ms-input-placeholder {
  color: var(--accent-color);
}

.contact-me-dialog__body {
  display: block;
  align-self: center;
  height: 10rem;
  margin: 0.5rem auto;
  width: 90%;
  border-radius: 0.2rem;
  border: 1px solid rgba(0, 0, 0, 0.15);
  font: 0.9em sans-serif;
  padding: 0.5rem;
  resize: none;
  color: var(--text-color);
  background: var(--background-color);
  border: 1px solid var(--text-color);
}

.contact-me-dialog__submit-btn {
  color: #fff !important;
  margin: 1rem !important;
  margin-bottom: 2.5rem !important;
  background: var(--accent-color) !important;
  width: 6rem !important;
  border: none !important;
  font-weight: 600;
  cursor: pointer;
}

.contact-me-dialog__submit-btn:disabled {
  filter: grayscale(50%);
  cursor: progress;
}

.contact-me-dialog-closer {
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0px 0px 3px 2px var(--shadow-color);
  border-radius: 50%;
  right: 1rem;
  top: 1rem;
  height: 1.5rem;
  width: 1.5rem;
  cursor: pointer;
  color: rgb(255, 91, 91);
}

.contact-me-dialog__status {
  color: var(--accent-color);
  margin: 1rem;
  font-size: 7rem;
}

.grow-enter-active,
.grow-leave-active {
  transition: all 225ms ease-out;
}

.grow-enter,
.grow-leave-to {
  opacity: 0;
  transform: scale(0.4);
}

.modal-leave,
.modal-enter-to {
  opacity: 1;
  transform: scale(1);
}
</style>
