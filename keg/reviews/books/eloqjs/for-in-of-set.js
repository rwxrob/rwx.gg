let devices = new Set(["Linux", "Mac", "Windows"]);
devices["category"] = "desktops";

for (let device in devices) {
   console.log(device); // "category"
}

for (let device of devices) {
    console.log(device); // "Linux", "Mac", "Windows"
}
