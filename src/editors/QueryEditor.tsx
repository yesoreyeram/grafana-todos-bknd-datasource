import defaults from 'lodash/defaults';

import React from 'react';
import { Select } from '@grafana/ui';
import { QueryEditorProps } from '@grafana/data';
import { DataSource } from '../DataSource';
import { defaultQuery, MyDataSourceOptions, MyQuery, EntityType } from '../types';
import { DummyEditor } from './Dummy.QueryEditor';
import { TodosEditor } from './Todos.QueryEditor';
import { JSONPlaceholderEditor } from './JSONPlaceholder.QueryEditor';

const EntityTypes = [
  { label: 'Dummy', value: EntityType.Dummy },
  { label: 'Todos', value: EntityType.Todos },
  { label: 'JSON Placeholder', value: EntityType.JSONPlaceholder },
];

type Props = QueryEditorProps<DataSource, MyQuery, MyDataSourceOptions>;

export const QueryEditor: React.FC<Props> = props => {
  let { query, onChange, onRunQuery } = props;

  query = defaults(query, defaultQuery);

  const onEntityTypeChange = (newEntity: EntityType) => {
    onChange({ ...query, entityType: newEntity });
    onRunQuery();
  };

  return (
    <div className="gf-form-inline">
      <div className="gf-form">
        <label className="gf-form-label query-keyword width-8">Entity</label>
        <Select
          className="width-12 min-width-12"
          options={EntityTypes}
          value={query.entityType}
          onChange={e => onEntityTypeChange(e.value as EntityType)}
        />
        {query.entityType === EntityType.Dummy && (
          <DummyEditor query={query} onChange={onChange} onRunQuery={onRunQuery}></DummyEditor>
        )}
        {query.entityType === EntityType.Todos && (
          <TodosEditor query={query} onChange={onChange} onRunQuery={onRunQuery}></TodosEditor>
        )}
        {query.entityType === EntityType.JSONPlaceholder && (
          <JSONPlaceholderEditor query={query} onChange={onChange} onRunQuery={onRunQuery}></JSONPlaceholderEditor>
        )}
      </div>
    </div>
  );
};
