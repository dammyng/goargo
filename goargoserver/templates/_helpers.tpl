{{- define "goargoserver.name" -}}
  {{- default "goargoserver" .Chart.Name }}
{{- end }}

{{- define "goargoserver.fullname" -}}
  {{- printf "%s-%s" .Release.Name (include "goargoserver.name" .) }}
{{- end }}
