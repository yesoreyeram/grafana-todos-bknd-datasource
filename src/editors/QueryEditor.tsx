import defaults from 'lodash/defaults';

import React from 'react';
import { Select } from '@grafana/ui';
import { QueryEditorProps } from '@grafana/data';
import { DataSource } from '../DataSource';
import { defaultQuery, MyDataSourceOptions, MyQuery, EntitiyType } from '../types';
import { DummyEditor } from './Dummy.QueryEditor';
import { TodosEditor } from './Todos.QueryEditor';

const EntityTypes = [
  { label: 'Dummy', value: EntitiyType.Dummy },
  { label: 'Todos', value: EntitiyType.Todos },
];

type Props = QueryEditorProps<DataSource, MyQuery, MyDataSourceOptions>;

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
