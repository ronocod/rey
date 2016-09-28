package com.ronocod.gotos;

import android.support.v7.widget.RecyclerView;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.CheckBox;
import android.widget.TextView;

import go.teanga.Store;
import go.teanga.Subscriber;
import go.teanga.Todo;

class TodoAdapter extends RecyclerView.Adapter<TodoAdapter.ViewHolder> implements Subscriber {

    private final Store store;

    TodoAdapter(Store store) {
        this.store = store;
    }

    @Override
    public ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        LayoutInflater inflater = LayoutInflater.from(parent.getContext());
        return new ViewHolder(inflater.inflate(R.layout.todo_list_content, parent, false));
    }

    @Override
    public void onBindViewHolder(final ViewHolder holder, int position) {
        Todo todo = store.getState().itemAtIndex(position);
        holder.textView.setText(todo.getName());
        holder.checkBox.setChecked(todo.getDone());

        holder.checkBox.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                int newPosition = holder.getAdapterPosition();
                Todo todo = store.getState().itemAtIndex(newPosition);
                if (todo.getDone()) {
                    store.uncheckItemAtIndex(newPosition);
                } else {
                    store.checkItemAtIndex(newPosition);
                }
            }
        });
    }

    @Override
    public int getItemCount() {
        return (int) store.getState().count();
    }

    @Override public void update() {
        Log.d("LoL", "Those bants hi");
        notifyDataSetChanged();
    }

    class ViewHolder extends RecyclerView.ViewHolder {
        final TextView textView;
        final CheckBox checkBox;

        ViewHolder(View view) {
            super(view);
            textView = (TextView) view.findViewById(R.id.item_text);
            checkBox = (CheckBox) view.findViewById(R.id.item_check);
        }

    }
}
