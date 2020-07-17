import { useEffect, useState } from "react";

const useNotification = (): any => {
    const [subscription, setSubscription] = useState<PushSubscriptionJSON>()

    useEffect(() => {

        const getSubscription = async () => {
            const permissionRequest = await Notification.requestPermission()
            if(permissionRequest === "granted") {
                const registration = await navigator.serviceWorker.register("/service-worker.js");
                const subscriptionOptions: PushSubscriptionOptionsInit = {
                    applicationServerKey: "BKZz5Ou266l-l6w5ZuVSqLZVzWmcNr3MikxSSX6hw4_QQb5wAozT9_xVmHJbmLHACnAkeJDUNTvwVjZWDjYXCN8",
                    userVisibleOnly: true
                } 
                const subscription = await registration.pushManager.subscribe(subscriptionOptions)
                setSubscription(subscription.toJSON());
            }
        }

        getSubscription();

    }, []);

    return subscription;
}

export {
    useNotification
}