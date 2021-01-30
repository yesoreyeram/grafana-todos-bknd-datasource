import React, { ChangeEvent } from 'react';
import { LegacyForms } from '@grafana/ui';
import { MyQuery } from '../types';

const { FormField } = LegacyForms;

interface JSONEditorProps {
  query: MyQuery;
  onChange: (value: MyQuery) => void;
  onRunQuery: () => void;
}

export const JSONQueryEditor: React.FC<JSONEditorProps> = props => {
  let { onChange, query, onRunQuery } = props;

  const onJSONURLChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange({ ...query, jsonURL: event.target.value });
    onRunQuery();
  };

  return (
    <>
      <FormField
        labelWidth={8}
        value={props.query.jsonURL}
        onChange={onJSONURLChange}
        label="JSON URL"
        tooltip="JSON URL"
      />
    </>
  );
};
