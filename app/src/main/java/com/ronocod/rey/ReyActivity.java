package com.ronocod.rey;

import android.os.AsyncTask;
import android.os.Bundle;
import android.support.design.widget.FloatingActionButton;
import android.support.v4.widget.SwipeRefreshLayout;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.view.View;
import android.widget.TextView;

import java.util.HashMap;

import go.rey.Person;
import go.rey.Rey;
import go.rey.State;
import go.rey.Store;
import go.rey.Subscriber;
import io.intercom.android.sdk.Intercom;

public class ReyActivity extends AppCompatActivity implements Subscriber {

    private Store store;
    private TextView textView;
    private SwipeRefreshLayout refreshLayout;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Intercom.client().registerUnidentifiedUser();

        store = Rey.newStore();

        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);
        toolbar.setTitle(getTitle());

        FloatingActionButton fab = (FloatingActionButton) findViewById(R.id.fab);
        fab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                fetchNext();
            }
        });

        store.setSubscriber(this);
        textView = (TextView) findViewById(R.id.item_text);
        refreshLayout = ((SwipeRefreshLayout) findViewById(R.id.refreshLayout));

        fetchNext();
    }

    private void fetchNext() {
        AsyncTask.execute(new Runnable() {
            @Override public void run() {
                store.fetchNextPerson();
            }
        });
    }

    @Override public void update() {
        final State state = store.getState();
        HashMap<String, Object> attributes = new HashMap<>();
        attributes.put("has_requested", true);
        Intercom.client().updateUser(attributes);
        Intercom.client().logEvent("request");

        final Person person = state.getCurrentPerson();
        final String text = person == null
                ? "No-one loaded yet"
                : person.toString()
                .replace("{", "{\n\n")
                .replace(",", ",\n")
                .replace("}", "\n}");
        final boolean isFetching = state.getIsFetching();

        runOnUiThread(new Runnable() {
            @Override public void run() {
                refreshLayout.setRefreshing(isFetching);
                textView.setText(text);
            }
        });
    }
}
