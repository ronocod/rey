package com.ronocod.rey;

import android.app.Application;

import io.intercom.android.sdk.Intercom;
import io.intercom.android.sdk.logger.IntercomLogger;

public class ReyApplication extends Application {

    @Override public void onCreate() {
        super.onCreate();

        Intercom.initialize(this, "android_sdk-a318b76d4d43dc331ca845e57516de1f6ae6ca72", "bn9as3lt");
        Intercom.setLogLevel(IntercomLogger.VERBOSE);

    }
}
