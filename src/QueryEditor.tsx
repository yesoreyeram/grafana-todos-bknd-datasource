import defaults from 'lodash/defaults';

import React, { ChangeEvent } from 'react';
import { LegacyForms, Select, Checkbox } from '@grafana/ui';
import { QueryEditorProps } from '@grafana/data';
import { DataSource } from './DataSource';
import { defaultQuery, MyDataSourceOptions, MyQuery, EntitiyType } from './types';

const { FormField } = LegacyForms;

const EntityTypes = [
  { label: 'Dummy', value: EntitiyType.Dummy },
  { label: 'Todos', value: EntitiyType.Todos },
];

type Props = QueryEditorProps<DataSource, MyQuery, MyDataSourceOptions>;

interface DummyProps {
  query: MyQuery;
  onChange: (value: MyQuery) => void;
  onRunQuery: () => void;
}

const DummyEditor: React.FC<DummyProps> = props => {
  let { onChange, query, onRunQuery } = props;

  const onQueryTextChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, queryText: event.target.value });
    onRunQuery();
  };

  const onConstantChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, constant: parseFloat(event.target.value) });
    onRunQuery();
  };

  return (
    <>
      <FormField
        width={4}
        value={props.query.constant}
        onChange={onConstantChange}
        label="Constant"
        type="number"
        step="0.1"
      />
      <FormField
        labelWidth={8}
        value={props.query.queryText}
        onChange={onQueryTextChange}
        label="Query Text"
        tooltip="Not used yet"
      />
    </>
  );
};

interface TodoEditorProps {
  query: MyQuery;
  onChange: (value: MyQuery) => void;
  onRunQuery: () => void;
}

const TodosEditor: React.FC<TodoEditorProps> = props => {
  let { onChange, query, onRunQuery } = props;

  const onHideFinishedTodosChange = () => {
    onChange({ ...query, hideFinishedTodos: !query.hideFinishedTodos });
    onRunQuery();
  };

  const onNumberofTodosChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, numberOfTodos: parseInt(event.target.value) });
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
      />
      <label className="gf-form-label query-keyword width-8">Hide Finished Todos</label>
      <Checkbox css={{}} value={props.query.hideFinishedTodos} onChange={e => onHideFinishedTodosChange()}></Checkbox>
    </>
  );
};

export const QueryEditor: React.FC<Props> = props => {
  let { query, onChange, onRunQuery } = props;

  query = defaults(query, defaultQuery);

  const onEntityTypeChange = (newEntity: EntitiyType) => {
    onChange({ ...query, entityType: newEntity });
    onRunQuery();
  };

  return (
    <div className="gf-form-inline">
      <div className="gf-form">
        <label className="gf-form-label query-keyword width-8">Entity</label>
        <Select
          options={EntityTypes}
          value={query.entityType}
          onChange={e => onEntityTypeChange(e.value as EntitiyType)}
        />
        {query.entityType === EntitiyType.Dummy && (
          <DummyEditor query={query} onChange={onChange} onRunQuery={onRunQuery}></DummyEditor>
        )}
        {query.entityType === EntitiyType.Todos && (
          <TodosEditor query={query} onChange={onChange} onRunQuery={onRunQuery}></TodosEditor>
        )}
      </div>
    </div>
  );
};
