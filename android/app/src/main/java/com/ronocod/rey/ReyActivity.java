package com.ronocod.rey;

import android.os.Bundle;
import android.support.design.widget.FloatingActionButton;
import android.support.v4.widget.SwipeRefreshLayout;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.view.View;
import android.widget.TextView;

import com.ronocod.rey.core.Core;
import com.ronocod.rey.core.Person;
import com.ronocod.rey.core.State;
import com.ronocod.rey.core.Store;
import com.ronocod.rey.core.Subscriber;

public class ReyActivity extends AppCompatActivity implements Subscriber {

    private Store store;
    private TextView textView;
    private SwipeRefreshLayout refreshLayout;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        store = Core.newStore();

        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);
        toolbar.setTitle(getTitle());

        FloatingActionButton fab = (FloatingActionButton) findViewById(R.id.fab);
        fab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                Core.fetchNextPerson(store);
            }
        });

        textView = (TextView) findViewById(R.id.item_text);
        refreshLayout = (SwipeRefreshLayout) findViewById(R.id.refreshLayout);

        store.subscribe(this);
        Core.fetchNextPerson(store);
    }

    @Override protected void onDestroy() {
        store.unsubscribe(this);
        super.onDestroy();
    }

    @Override public void update(State state) {
        Person person = state.getCurrentPerson();
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
