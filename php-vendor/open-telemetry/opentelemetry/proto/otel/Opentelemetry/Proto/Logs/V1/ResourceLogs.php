<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: opentelemetry/proto/logs/v1/logs.proto

namespace Opentelemetry\Proto\Logs\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * A collection of ScopeLogs from a Resource.
 *
 * Generated from protobuf message <code>opentelemetry.proto.logs.v1.ResourceLogs</code>
 */
class ResourceLogs extends \Google\Protobuf\Internal\Message
{
    /**
     * The resource for the logs in this message.
     * If this field is not set then resource info is unknown.
     *
     * Generated from protobuf field <code>.opentelemetry.proto.resource.v1.Resource resource = 1;</code>
     */
    protected $resource = null;
    /**
     * A list of ScopeLogs that originate from a resource.
     *
     * Generated from protobuf field <code>repeated .opentelemetry.proto.logs.v1.ScopeLogs scope_logs = 2;</code>
     */
    private $scope_logs;
    /**
     * A list of InstrumentationLibraryLogs that originate from a resource.
     * This field is deprecated and will be removed after grace period expires on June 15, 2022.
     * During the grace period the following rules SHOULD be followed:
     * For Binary Protobufs
     * ====================
     * Binary Protobuf senders SHOULD NOT set instrumentation_library_logs. Instead
     * scope_logs SHOULD be set.
     * Binary Protobuf receivers SHOULD check if instrumentation_library_logs is set
     * and scope_logs is not set then the value in instrumentation_library_logs
     * SHOULD be used instead by converting InstrumentationLibraryLogs into ScopeLogs.
     * If scope_logs is set then instrumentation_library_logs SHOULD be ignored.
     * For JSON
     * ========
     * JSON senders that set instrumentation_library_logs field MAY also set
     * scope_logs to carry the same logs, essentially double-publishing the same data.
     * Such double-publishing MAY be controlled by a user-settable option.
     * If double-publishing is not used then the senders SHOULD set scope_logs and
     * SHOULD NOT set instrumentation_library_logs.
     * JSON receivers SHOULD check if instrumentation_library_logs is set and
     * scope_logs is not set then the value in instrumentation_library_logs
     * SHOULD be used instead by converting InstrumentationLibraryLogs into ScopeLogs.
     * If scope_logs is set then instrumentation_library_logs field SHOULD be ignored.
     *
     * Generated from protobuf field <code>repeated .opentelemetry.proto.logs.v1.InstrumentationLibraryLogs instrumentation_library_logs = 1000 [deprecated = true];</code>
     * @deprecated
     */
    private $instrumentation_library_logs;
    /**
     * This schema_url applies to the data in the "resource" field. It does not apply
     * to the data in the "scope_logs" field which have their own schema_url field.
     *
     * Generated from protobuf field <code>string schema_url = 3;</code>
     */
    protected $schema_url = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Opentelemetry\Proto\Resource\V1\Resource $resource
     *           The resource for the logs in this message.
     *           If this field is not set then resource info is unknown.
     *     @type \Opentelemetry\Proto\Logs\V1\ScopeLogs[]|\Google\Protobuf\Internal\RepeatedField $scope_logs
     *           A list of ScopeLogs that originate from a resource.
     *     @type \Opentelemetry\Proto\Logs\V1\InstrumentationLibraryLogs[]|\Google\Protobuf\Internal\RepeatedField $instrumentation_library_logs
     *           A list of InstrumentationLibraryLogs that originate from a resource.
     *           This field is deprecated and will be removed after grace period expires on June 15, 2022.
     *           During the grace period the following rules SHOULD be followed:
     *           For Binary Protobufs
     *           ====================
     *           Binary Protobuf senders SHOULD NOT set instrumentation_library_logs. Instead
     *           scope_logs SHOULD be set.
     *           Binary Protobuf receivers SHOULD check if instrumentation_library_logs is set
     *           and scope_logs is not set then the value in instrumentation_library_logs
     *           SHOULD be used instead by converting InstrumentationLibraryLogs into ScopeLogs.
     *           If scope_logs is set then instrumentation_library_logs SHOULD be ignored.
     *           For JSON
     *           ========
     *           JSON senders that set instrumentation_library_logs field MAY also set
     *           scope_logs to carry the same logs, essentially double-publishing the same data.
     *           Such double-publishing MAY be controlled by a user-settable option.
     *           If double-publishing is not used then the senders SHOULD set scope_logs and
     *           SHOULD NOT set instrumentation_library_logs.
     *           JSON receivers SHOULD check if instrumentation_library_logs is set and
     *           scope_logs is not set then the value in instrumentation_library_logs
     *           SHOULD be used instead by converting InstrumentationLibraryLogs into ScopeLogs.
     *           If scope_logs is set then instrumentation_library_logs field SHOULD be ignored.
     *     @type string $schema_url
     *           This schema_url applies to the data in the "resource" field. It does not apply
     *           to the data in the "scope_logs" field which have their own schema_url field.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Opentelemetry\Proto\Logs\V1\Logs::initOnce();
        parent::__construct($data);
    }

    /**
     * The resource for the logs in this message.
     * If this field is not set then resource info is unknown.
     *
     * Generated from protobuf field <code>.opentelemetry.proto.resource.v1.Resource resource = 1;</code>
     * @return \Opentelemetry\Proto\Resource\V1\Resource|null
     */
    public function getResource()
    {
        return $this->resource;
    }

    public function hasResource()
    {
        return isset($this->resource);
    }

    public function clearResource()
    {
        unset($this->resource);
    }

    /**
     * The resource for the logs in this message.
     * If this field is not set then resource info is unknown.
     *
     * Generated from protobuf field <code>.opentelemetry.proto.resource.v1.Resource resource = 1;</code>
     * @param \Opentelemetry\Proto\Resource\V1\Resource $var
     * @return $this
     */
    public function setResource($var)
    {
        GPBUtil::checkMessage($var, \Opentelemetry\Proto\Resource\V1\Resource::class);
        $this->resource = $var;

        return $this;
    }

    /**
     * A list of ScopeLogs that originate from a resource.
     *
     * Generated from protobuf field <code>repeated .opentelemetry.proto.logs.v1.ScopeLogs scope_logs = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getScopeLogs()
    {
        return $this->scope_logs;
    }

    /**
     * A list of ScopeLogs that originate from a resource.
     *
     * Generated from protobuf field <code>repeated .opentelemetry.proto.logs.v1.ScopeLogs scope_logs = 2;</code>
     * @param \Opentelemetry\Proto\Logs\V1\ScopeLogs[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setScopeLogs($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Opentelemetry\Proto\Logs\V1\ScopeLogs::class);
        $this->scope_logs = $arr;

        return $this;
    }

    /**
     * A list of InstrumentationLibraryLogs that originate from a resource.
     * This field is deprecated and will be removed after grace period expires on June 15, 2022.
     * During the grace period the following rules SHOULD be followed:
     * For Binary Protobufs
     * ====================
     * Binary Protobuf senders SHOULD NOT set instrumentation_library_logs. Instead
     * scope_logs SHOULD be set.
     * Binary Protobuf receivers SHOULD check if instrumentation_library_logs is set
     * and scope_logs is not set then the value in instrumentation_library_logs
     * SHOULD be used instead by converting InstrumentationLibraryLogs into ScopeLogs.
     * If scope_logs is set then instrumentation_library_logs SHOULD be ignored.
     * For JSON
     * ========
     * JSON senders that set instrumentation_library_logs field MAY also set
     * scope_logs to carry the same logs, essentially double-publishing the same data.
     * Such double-publishing MAY be controlled by a user-settable option.
     * If double-publishing is not used then the senders SHOULD set scope_logs and
     * SHOULD NOT set instrumentation_library_logs.
     * JSON receivers SHOULD check if instrumentation_library_logs is set and
     * scope_logs is not set then the value in instrumentation_library_logs
     * SHOULD be used instead by converting InstrumentationLibraryLogs into ScopeLogs.
     * If scope_logs is set then instrumentation_library_logs field SHOULD be ignored.
     *
     * Generated from protobuf field <code>repeated .opentelemetry.proto.logs.v1.InstrumentationLibraryLogs instrumentation_library_logs = 1000 [deprecated = true];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     * @deprecated
     */
    public function getInstrumentationLibraryLogs()
    {
        @trigger_error('instrumentation_library_logs is deprecated.', E_USER_DEPRECATED);
        return $this->instrumentation_library_logs;
    }

    /**
     * A list of InstrumentationLibraryLogs that originate from a resource.
     * This field is deprecated and will be removed after grace period expires on June 15, 2022.
     * During the grace period the following rules SHOULD be followed:
     * For Binary Protobufs
     * ====================
     * Binary Protobuf senders SHOULD NOT set instrumentation_library_logs. Instead
     * scope_logs SHOULD be set.
     * Binary Protobuf receivers SHOULD check if instrumentation_library_logs is set
     * and scope_logs is not set then the value in instrumentation_library_logs
     * SHOULD be used instead by converting InstrumentationLibraryLogs into ScopeLogs.
     * If scope_logs is set then instrumentation_library_logs SHOULD be ignored.
     * For JSON
     * ========
     * JSON senders that set instrumentation_library_logs field MAY also set
     * scope_logs to carry the same logs, essentially double-publishing the same data.
     * Such double-publishing MAY be controlled by a user-settable option.
     * If double-publishing is not used then the senders SHOULD set scope_logs and
     * SHOULD NOT set instrumentation_library_logs.
     * JSON receivers SHOULD check if instrumentation_library_logs is set and
     * scope_logs is not set then the value in instrumentation_library_logs
     * SHOULD be used instead by converting InstrumentationLibraryLogs into ScopeLogs.
     * If scope_logs is set then instrumentation_library_logs field SHOULD be ignored.
     *
     * Generated from protobuf field <code>repeated .opentelemetry.proto.logs.v1.InstrumentationLibraryLogs instrumentation_library_logs = 1000 [deprecated = true];</code>
     * @param \Opentelemetry\Proto\Logs\V1\InstrumentationLibraryLogs[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     * @deprecated
     */
    public function setInstrumentationLibraryLogs($var)
    {
        @trigger_error('instrumentation_library_logs is deprecated.', E_USER_DEPRECATED);
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Opentelemetry\Proto\Logs\V1\InstrumentationLibraryLogs::class);
        $this->instrumentation_library_logs = $arr;

        return $this;
    }

    /**
     * This schema_url applies to the data in the "resource" field. It does not apply
     * to the data in the "scope_logs" field which have their own schema_url field.
     *
     * Generated from protobuf field <code>string schema_url = 3;</code>
     * @return string
     */
    public function getSchemaUrl()
    {
        return $this->schema_url;
    }

    /**
     * This schema_url applies to the data in the "resource" field. It does not apply
     * to the data in the "scope_logs" field which have their own schema_url field.
     *
     * Generated from protobuf field <code>string schema_url = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setSchemaUrl($var)
    {
        GPBUtil::checkString($var, True);
        $this->schema_url = $var;

        return $this;
    }

}

