{{/*
Common name helpers for the kuberik chart.
*/}}

{{- define "kuberik.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "kuberik.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name (include "kuberik.name" .) | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}

{{- define "kuberik.labels" -}}
app.kubernetes.io/name: {{ include "kuberik.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
helm.sh/chart: {{ printf "%s-%s" .Chart.Name .Chart.Version | quote }}
{{- end -}}

{{- define "kuberik.rolloutControllerImage" -}}
{{- $tag := .Values.rolloutController.image.tag -}}
{{- if not $tag -}}
{{- $tag = printf "v%s" .Chart.AppVersion -}}
{{- end -}}
{{ printf "%s:%s" .Values.rolloutController.image.repository $tag }}
{{- end -}}
