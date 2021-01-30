import React, { ChangeEvent } from 'react';
import { LegacyForms, Checkbox } from '@grafana/ui';
import { MyQuery } from '../types';

const { FormField } = LegacyForms;

interface TodoEditorProps {
  query: MyQuery;
  onChange: (value: MyQuery) => void;
  onRunQuery: () => void;
}

export const TodosEditor: React.FC<TodoEditorProps> = props => {
  let { onChange, query, onRunQuery } = props;

  const onHideFinishedTodosChange = () => {
    onChange({ ...query, hideFinishedTodos: !query.hideFinishedTodos });
    onRunQuery();
  };

  const onNumberofTodosChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, numberOfTodos: parseInt(event.target.value, 10) });
    onRunQuery();
  };

  return (
    <>
      <FormField
        width={4}
        labelWidth={8}
        value={props.query.numberOfTodos}
        onChange={onNumberofTodosChange}
        label="Number of Todos"
        type="number"
        step="1"
        placeholder="200"
        required
        min={0}
        max={200}
      />
      <label className="gf-form-label query-keyword width-8">Hide Finished Todos</label>
      <Checkbox css={{}} value={props.query.hideFinishedTodos} onChange={e => onHideFinishedTodosChange()}></Checkbox>
    </>
  );
};
