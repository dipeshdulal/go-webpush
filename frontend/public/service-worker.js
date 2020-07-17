this.addEventListener("push", (event) => {
  if (event.data) {
    const notifPromise = this.registration.showNotification(event.data.text());
    event.waitUntil(notifPromise);
  } else {
    console.log("No data in this push event.");
  }
});

this.addEventListener("install", (event) => {
    this.skipWaiting();
})
