import React, { ChangeEvent } from 'react';
import { LegacyForms } from '@grafana/ui';
import { MyQuery } from '../types';

const { FormField } = LegacyForms;

interface DummyProps {
  query: MyQuery;
  onChange: (value: MyQuery) => void;
  onRunQuery: () => void;
}

export const DummyEditor: React.FC<DummyProps> = props => {
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
