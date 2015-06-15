kd             = require 'kd'
JView          = require 'app/jview'
remote         = require('app/remote').getInstance()
KDButtonView   = kd.ButtonView
KDListItemView = kd.ListItemView

DUMMY_DATA     =
  setupData    :
    name       : 'Airbrake'
    logo       : 'https://koding-cdn.s3.amazonaws.com/temp-images/airbrake.png'
    summary    : 'Error monitoring and handling.'
    desc       : """
      <p>Airbrake captures errors generated by your team's web applications, and aggregates the results for easy review.</p>
      <p>This integration will post notifications in Slack when an error is reported in Airbrake.</p>
    """
    channels   : [
      { name   : '#general',   id: '123' }
      { name   : '#random',    id: '124' }
      { name   : '#off-topic', id: '125' }
    ]
  detailsData  :
    name    : 'Airbrake'
    logo    : 'https://koding-cdn.s3.amazonaws.com/temp-images/airbrake.png'
    summary : 'Error monitoring and handling.'
    desc    : """
      <p>Airbrake captures errors generated by your team's web applications, and aggregates the results for easy review.</p>
      <p>This integration will post notifications in Slack when an error is reported in Airbrake.</p>
    """
    channels : [
      { name: '#general',   id: '123' }
      { name: '#random',    id: '124' }
      { name: '#off-topic', id: '125' }
    ]
    selectedChannel : '123'
    webhookUrl      : 'https://hooks.koding.com/services/123/abc/def'
    instructions    : """
      #### Step 1

      In your Airbrake dashboard, click on the menu button in the top navigation, and then select **Integrations** from the sub-menu.

      ![airbrake_step1.png](https://koding-cdn.s3.amazonaws.com/temp-images/airbrake_step1.png)


      #### Step 2

      Click on Slack in the list of integrations, add `https://hooks.slack.com/services/123/abc/def` as the Slack webhook, and choose your notification settings. Press the Save button when you're done.

      ![airbrake_step2.png](https://koding-cdn.s3.amazonaws.com/temp-images/airbrake_step2.png)
    """


module.exports = class AdminIntegrationItemView extends KDListItemView

  JView.mixin @prototype

  constructor: (options = {}, data) ->

    options.cssClass = 'integration-view'

    super options, data

    { @integrationType } = @getData()

    @button    = new KDButtonView
      cssClass : 'solid compact green add'
      title    : if @integrationType is 'new' then 'Add' else 'Configure'
      loader   : yes
      callback : =>
        if      @integrationType is 'new' then @fetchIntegrationChannels()
        else if @integrationType is 'configured'
          @fetchIntegrationDetails()


  fetchIntegrationChannels: ->
    data = @getData()

    @fetchChannels (err, channels) =>
      return  if err

      data.channels = channels

      @button.hideLoader()
      @emit 'IntegrationGroupsFetched', data


  fetchIntegrationDetails: ->
    { title, summary, description, iconPath } = @getData()

    remote.api.JAccount.some {}, {}, (err, data) =>
      @button.hideLoader()
      @emit 'IntegrationConfigureRequested', DUMMY_DATA.detailsData


  fetchChannels: (callback) ->
    kd.singletons.socialapi.account.fetchChannels (err, channels) =>
      return callback err  if err

      decoratedChannels = []
      for channel in channels
        { id, typeConstant, name, purpose, participantsPreview } = channel

        # TODO after refactoring the private channels, we also need
        # to add them here
        if typeConstant is 'topic' or typeConstant is 'group'
          decoratedChannels.push { name:"##{name}", id }

      callback null, decoratedChannels

  pistachio: ->

    { title, summary, iconPath } = @getData()

    return """
      <img src="#{iconPath}" />
      {p{ #(title)}}
      {span{ #(summary)}}
      {{> @button}}
    """