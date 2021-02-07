import React, { ChangeEvent } from 'react';
import { LegacyForms } from '@grafana/ui';
import { DataSourcePluginOptionsEditorProps } from '@grafana/data';
import { MyDataSourceOptions, MySecureJsonData } from '../types';

const { SecretFormField, FormField } = LegacyForms;

interface Props extends DataSourcePluginOptionsEditorProps<MyDataSourceOptions> {}

export const ConfigEditor: React.FC<Props> = props => {
  const { onOptionsChange, options } = props;
  const { jsonData, secureJsonFields } = options;
  const secureJsonData = (options.secureJsonData || {}) as MySecureJsonData;

  const onValueChange = (event: ChangeEvent<HTMLInputElement>) => {
    const jsonData = {
      ...options.jsonData,
      promValue: parseInt(event.target.value, 10),
    };
    onOptionsChange({ ...options, jsonData });
  };

  const onPathChange = (event: ChangeEvent<HTMLInputElement>) => {
    const jsonData = {
      ...options.jsonData,
      path: event.target.value,
    };
    onOptionsChange({ ...options, jsonData });
  };

  const onDefaultJSONChange = (event: ChangeEvent<HTMLInputElement>) => {
    const jsonData = {
      ...options.jsonData,
      defaultJSONURL: event.target.value,
    };
    onOptionsChange({ ...options, jsonData });
  };

  const onAPIKeyChange = (event: ChangeEvent<HTMLInputElement>) => {
    onOptionsChange({
      ...options,
      secureJsonData: {
        apiKey: event.target.value,
      },
    });
  };

  const onResetAPIKey = () => {
    onOptionsChange({
      ...options,
      secureJsonFields: {
        ...options.secureJsonFields,
        apiKey: false,
      },
      secureJsonData: {
        ...options.secureJsonData,
        apiKey: '',
      },
    });
  };

  return (
    <div className="gf-form-group">
      <div className="gf-form">
        <FormField
          label="Path"
          labelWidth={6}
          inputWidth={20}
          onChange={onValueChange}
          value={jsonData.promValue || 10}
          placeholder="Sample value in prometheus"
        />
      </div>
      <div className="gf-form">
        <FormField
          label="Path"
          labelWidth={6}
          inputWidth={20}
          onChange={onPathChange}
          value={jsonData.path || ''}
          placeholder="json field returned to frontend"
        />
      </div>
      <div className="gf-form">
        <FormField
          label="Path"
          labelWidth={6}
          inputWidth={20}
          onChange={onDefaultJSONChange}
          value={jsonData.defaultJSONURL || ''}
          placeholder="Default JSON URL"
        />
      </div>
      <div className="gf-form-inline">
        <div className="gf-form">
          <SecretFormField
            isConfigured={(secureJsonFields && secureJsonFields.apiKey) as boolean}
            value={secureJsonData.apiKey || ''}
            label="API Key"
            placeholder="secure json field (backend only)"
            labelWidth={6}
            inputWidth={20}
            onReset={onResetAPIKey}
            onChange={onAPIKeyChange}
          />
        </div>
      </div>
    </div>
  );
};
